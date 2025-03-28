/*
Copyright 2021 Flant JSC
Licensed under the Deckhouse Platform Enterprise Edition (EE) license. See https://github.com/deckhouse/deckhouse/blob/main/ee/LICENSE
*/

package cache

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"slices"
	"strings"
	"sync"
	"time"

	k8sversion "k8s.io/apimachinery/pkg/version"
)

const (
	renewTokenPeriod = 30 * time.Second

	defaultTTL = 1 * time.Hour

	saPath    = "/var/run/secrets/kubernetes.io/serviceaccount/"
	caPath    = saPath + "ca.crt"
	tokenPath = saPath + "token"
	apiV1Path = "/api/v1"

	kubernetesAPIAddress = "https://kubernetes.default"
)

type Cache interface {
	Get(string, string) (bool, error)
	GetPreferredVersion(group, resource string) (string, error)
	GetCoreResources() (CoreResourcesDict, error)
	Check() error
}

var _ Cache = (*NamespacedDiscoveryCache)(nil)
var ErrNotFound = errors.New("not found")

type cacheEntry struct {
	TTL     time.Duration
	AddTime time.Time
}

func newCacheEntry(addTime time.Time) *cacheEntry {
	return &cacheEntry{
		AddTime: addTime,
		TTL:     defaultTTL,
	}
}

type namespacedCacheEntry struct {
	*cacheEntry
	Data map[string]bool
}

func newNamespacedCacheEntry(addTime time.Time) *namespacedCacheEntry {
	return &namespacedCacheEntry{
		cacheEntry: newCacheEntry(addTime),
		Data:       make(map[string]bool),
	}
}

type preferredVersionCacheEntry struct {
	*cacheEntry
	Version string
}

func newPreferredVersionCacheEntry(addTime time.Time, version string) *preferredVersionCacheEntry {
	return &preferredVersionCacheEntry{
		cacheEntry: newCacheEntry(addTime),
		Version:    version,
	}
}

type CoreResourcesDict map[string]struct{}

type coreResourcesCache struct {
	*cacheEntry
	dict CoreResourcesDict
}

type NamespacedDiscoveryCache struct {
	logger *log.Logger

	client *http.Client

	mu   sync.RWMutex
	data map[string]*namespacedCacheEntry

	muPv              sync.RWMutex
	preferredVersions map[string]*preferredVersionCacheEntry

	muCr          sync.RWMutex
	coreResources *coreResourcesCache

	now func() time.Time

	kubernetesAPIAddress string
}

func NewNamespacedDiscoveryCache(logger *log.Logger) *NamespacedDiscoveryCache {
	c := &NamespacedDiscoveryCache{
		logger:            logger,
		data:              make(map[string]*namespacedCacheEntry),
		preferredVersions: make(map[string]*preferredVersionCacheEntry),
		coreResources:     new(coreResourcesCache),
		now:               time.Now,

		kubernetesAPIAddress: kubernetesAPIAddress,
	}
	c.initClient()
	return c
}

func (c *NamespacedDiscoveryCache) Check() error {
	return Retry(func() (bool, error) {
		req, err := http.NewRequest(http.MethodGet, c.kubernetesAPIAddress+"/version", nil)
		if err != nil {
			return false, fmt.Errorf("check Kubernetes API create request: %w", err)
		}

		if _, err := c.execRequest(req, "check API", nil); err != nil {
			return true, err
		}
		return false, nil
	})
}

func (c *NamespacedDiscoveryCache) initClient() {
	tlsConfig := &tls.Config{}

	contentCA, err := os.ReadFile(caPath)
	if err == nil {
		caPool, err := x509.SystemCertPool()
		if err != nil {
			panic(fmt.Errorf("cannot get system cert pool: %v", err))
		}

		caPool.AppendCertsFromPEM(contentCA)
		tlsConfig.RootCAs = caPool
	} else {
		c.logger.Printf("%v: not in pod?", err)
	}

	baseTransport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}

	c.client = &http.Client{Transport: wrapKubeTransport(baseTransport)}
}

func (c *NamespacedDiscoveryCache) renewCacheOnce(apiGroup string, req *http.Request) error {
	var groupedResp Response
	_, err := c.execRequest(req, "renew namespaced cache", &groupedResp)
	if err != nil {
		return err
	}

	cache := newNamespacedCacheEntry(c.now())
	for _, resource := range groupedResp.Resources {
		cache.Data[resource.Name] = resource.Namespaced
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[apiGroup] = cache
	return nil
}

func (c *NamespacedDiscoveryCache) renewCache(apiGroup string) error {
	path := apiV1Path
	if apiGroup != "v1" {
		path = "/apis/" + apiGroup
	}

	return Retry(func() (bool, error) {
		req, err := http.NewRequest(http.MethodGet, c.kubernetesAPIAddress+path, nil)
		if err != nil {
			return false, fmt.Errorf("renew cache prepare request: %w", err)
		}

		if err := c.renewCacheOnce(apiGroup, req); err != nil {
			return true, err
		}

		return false, nil
	})
}

func (c *NamespacedDiscoveryCache) getAvailableAPIGroupVerionsInDescendingOrder(apiGroup string) ([]string, error) {
	path := "/apis/" + apiGroup

	availableVersions := make([]string, 0)

	err := Retry(func() (bool, error) {
		req, err := http.NewRequest(http.MethodGet, c.kubernetesAPIAddress+path, nil)
		if err != nil {
			return false, fmt.Errorf("build request for available apigroup versions: %w", err)
		}

		var apiGroupVersions APIGroupResponse
		_, err = c.execRequest(req, "request available apigroup versions", &apiGroupVersions)
		if err != nil {
			return true, fmt.Errorf("request available apigroup verions: %w", err)
		}

		for _, v := range apiGroupVersions.Versions {
			availableVersions = append(availableVersions, v.Version)
		}

		return false, nil
	})

	slices.SortFunc(availableVersions, func(v1, v2 string) int {
		return -(k8sversion.CompareKubeAwareVersionStrings(v1, v2))
	})

	return availableVersions, err
}

func (c *NamespacedDiscoveryCache) requestPreferredVersion(group, resource string) (string, error) {
	preferredVersion := ""

	availableVersions, err := c.getAvailableAPIGroupVerionsInDescendingOrder(group)
	if err != nil {
		return "", fmt.Errorf("get available apigroup versions: %w", err)
	}

	for _, version := range availableVersions {
		path := fmt.Sprintf("/apis/%s/%s", group, version)
		if err := Retry(func() (bool, error) {
			req, err := http.NewRequest(http.MethodGet, c.kubernetesAPIAddress+path, nil)
			if err != nil {
				return false, fmt.Errorf("request %s %s/%s version build error: %w", resource, group, version, err)
			}

			var apiResourceList APIResourceList
			_, err = c.execRequest(req, "request list of resources", &apiResourceList)
			if err != nil {
				return true, fmt.Errorf("request list of resources: %w", err)
			}

			if apiResourceList.Has(resource) {
				preferredVersion = version
			}

			return false, nil
		}); err != nil {
			if errors.Is(err, ErrNotFound) {
				continue
			}
			return "", fmt.Errorf("get preferred version: %w", err)
		}

		if preferredVersion != "" {
			break
		}
	}

	if preferredVersion == "" {
		return "", fmt.Errorf("failed to discover preferred version for %s.%s", resource, group)
	}

	return preferredVersion, nil
}

// GetCoreResources returns a dict of currently available resources from the core apigroup
func (c *NamespacedDiscoveryCache) GetCoreResources() (CoreResourcesDict, error) {
	coreResources := c.getCoreResourcesFromCache()
	if len(coreResources) != 0 {
		return coreResources, nil
	}

	coreResources, err := c.requestCoreResources()
	if err != nil {
		return nil, err
	}

	c.muCr.Lock()
	c.coreResources = &coreResourcesCache{
		cacheEntry: newCacheEntry(c.now()),
		dict:       coreResources,
	}
	c.muCr.Unlock()

	return coreResources, nil
}

func (c *NamespacedDiscoveryCache) getCoreResourcesFromCache() CoreResourcesDict {
	c.muCr.RLock()
	defer c.muCr.RUnlock()
	if len(c.coreResources.dict) != 0 && !c.isEntryExpired(c.coreResources.cacheEntry) {
		return c.coreResources.dict
	}

	return nil
}

func getResourceNameBeforeSlash(resourceName string) string {
	return strings.Split(resourceName, "/")[0]
}

func (c *NamespacedDiscoveryCache) requestCoreResources() (CoreResourcesDict, error) {
	var coreResources CoreResourcesDict
	if err := Retry(func() (bool, error) {
		req, err := http.NewRequest(http.MethodGet, c.kubernetesAPIAddress+apiV1Path, nil)
		if err != nil {
			return false, fmt.Errorf("build request for core resources: %w", err)
		}

		var apiResourceList APIResourceList
		_, err = c.execRequest(req, "request list of core resources", &apiResourceList)
		if err != nil {
			return true, fmt.Errorf("request list of core resources: %w", err)
		}

		discoveredCoreResources := make(CoreResourcesDict, len(apiResourceList.Resources))
		for _, resource := range apiResourceList.Resources {
			discoveredCoreResources[getResourceNameBeforeSlash(resource.Name)] = struct{}{}
		}

		coreResources = discoveredCoreResources

		return false, nil
	}); err != nil {
		return nil, fmt.Errorf("get list of core resources: %w", err)
	}

	return coreResources, nil
}

func (c *NamespacedDiscoveryCache) preferredVersionFromCache(group, resource string) string {
	c.muPv.RLock()
	defer c.muPv.RUnlock()

	entry, ok := c.preferredVersions[fmt.Sprintf("%s.%s", resource, group)]

	if ok && !c.isEntryExpired(entry.cacheEntry) {
		return entry.Version
	}

	return ""
}

func (c *NamespacedDiscoveryCache) GetPreferredVersion(group, resource string) (string, error) {
	version := c.preferredVersionFromCache(group, resource)
	if version != "" {
		return version, nil
	}

	version, err := c.requestPreferredVersion(group, resource)
	if err != nil {
		return "", err
	}

	c.muPv.Lock()
	defer c.muPv.Unlock()

	c.preferredVersions[fmt.Sprintf("%s.%s", resource, group)] = newPreferredVersionCacheEntry(c.now(), version)

	return version, nil
}

func (c *NamespacedDiscoveryCache) getFromCache(apiGroup string) (*namespacedCacheEntry, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.data[apiGroup]
	return entry, ok
}

func (c *NamespacedDiscoveryCache) Get(apiGroup, resource string) (bool, error) {
	namespacedInfo, ok := c.getFromCache(apiGroup)

	switch {
	case !ok:
		// there is no cache, renew
		if err := c.renewCache(apiGroup); err != nil {
			return false, err
		}

		namespacedInfo, _ = c.getFromCache(apiGroup)
	case c.isEntryExpired(namespacedInfo.cacheEntry):
		// cache is expired
		if err := c.renewCache(apiGroup); err != nil {
			// if there is an error, we could just use stale cache
			c.logger.Println(err)
		} else {
			namespacedInfo, _ = c.getFromCache(apiGroup)
		}
	}

	// if cache for api group exists but there is no resource, we should update the cache entry for the whole group
	namespaced, ok := namespacedInfo.Data[resource]
	if !ok {
		if err := c.renewCache(apiGroup); err != nil {
			return false, err
		}

		namespacedInfo, _ = c.getFromCache(apiGroup)

		namespaced, ok = namespacedInfo.Data[resource]
		if !ok {
			return false, fmt.Errorf("resource %s/%s is not found in cluster", apiGroup, resource)
		}
	}

	return namespaced, nil
}

func (c *NamespacedDiscoveryCache) isEntryExpired(e *cacheEntry) bool {
	return c.now().After(e.AddTime.Add(e.TTL))
}

func (c *NamespacedDiscoveryCache) execRequest(req *http.Request, logTag string, result interface{}) (string, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("%s: requesting error: %w", logTag, err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%s: decoding response error: %w", logTag, err)
	}

	if resp.StatusCode == 404 {
		return "", ErrNotFound
	}

	if resp.StatusCode/100 > 2 {
		return "", fmt.Errorf("%s: kube response error: %d %s", logTag, resp.StatusCode, respBody)
	}

	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return "", fmt.Errorf("%s: failed to unmarshal response: %w", logTag, err)
		}
	}

	return string(respBody), nil
}

// Resource is a single entry of the /apis/.../... endpoint response.
type Resource struct {
	Name       string `json:"name"`
	Namespaced bool   `json:"namespaced"`
}

// Response is a /apis/.../... endpoint response.
type Response struct {
	Resources []Resource `json:"resources"`
}

type PreferredVersion struct {
	// groupVersion specifies the API group and version in the form "group/version"
	GroupVersion string `json:"groupVersion"`
	// version specifies the version in the form of "version". This is to save
	// the clients the trouble of splitting the GroupVersion.
	Version string `json:"version"`
}

type APIGroupResponse struct {
	Versions []APIGroupVersion `json:"versions"`
}

type APIGroupVersion struct {
	Version string `json:"version"`
}

type APIResourceList struct {
	Resources []Resource `json:"resources"`
}

func (l *APIResourceList) Has(resource string) bool {
	for _, res := range l.Resources {
		if res.Name == resource {
			return true
		}
	}

	return false
}
