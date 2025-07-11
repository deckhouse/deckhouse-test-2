/*
Copyright 2021 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hooks

import (
	"errors"
	"fmt"
	"strings"

	"github.com/flant/addon-operator/pkg/module_manager/go_hook"
	"github.com/flant/addon-operator/sdk"
	"github.com/flant/shell-operator/pkg/kube_events_manager/types"
	"golang.org/x/exp/slog"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	sdkobjectpatch "github.com/deckhouse/module-sdk/pkg/object-patch"
)

var _ = sdk.RegisterFunc(&go_hook.HookConfig{
	Queue: "/modules/kube-proxy",
	Kubernetes: []go_hook.KubernetesConfig{
		{
			Name:              "kube_api_ep",
			ApiVersion:        "v1",
			Kind:              "Endpoints",
			NamespaceSelector: &types.NamespaceSelector{NameSelector: &types.NameSelector{MatchNames: []string{"default"}}},
			NameSelector:      &types.NameSelector{MatchNames: []string{"kubernetes"}},
			FilterFunc:        applyKubernetesAPIEndpointsFilter,
		},
	},
}, discoverAPIEndpointsHandler)

// KubernetesAPIEndpoints discovers kube api endpoints
type KubernetesAPIEndpoints struct {
	HostPort []string
}

func applyKubernetesAPIEndpointsFilter(obj *unstructured.Unstructured) (go_hook.FilterResult, error) {
	endpoint := &v1.Endpoints{}
	err := sdk.FromUnstructured(obj, endpoint)
	if err != nil {
		return nil, err
	}

	mh := &KubernetesAPIEndpoints{}

	for _, subset := range endpoint.Subsets {
		for _, address := range subset.Addresses {
			ip := address.IP
			for _, port := range subset.Ports {
				mh.HostPort = append(mh.HostPort, fmt.Sprintf("%s:%d", ip, port.Port))
			}
		}
	}

	return mh, nil
}

func discoverAPIEndpointsHandler(input *go_hook.HookInput) error {
	endpoints, err := sdkobjectpatch.UnmarshalToStruct[KubernetesAPIEndpoints](input.NewSnapshots, "kube_api_ep")
	if err != nil {
		return fmt.Errorf("failed to unmarshal kube_api_ep snapshot: %w", err)
	}

	if len(endpoints) == 0 {
		input.Logger.Error("kubernetes endpoints not found")
		return nil
	}

	fpp := endpoints[0]

	if len(fpp.HostPort) == 0 {
		return errors.New("no kubernetes apiserver endpoints host:port specified")
	}

	input.Logger.Info("cluster master addresses", slog.String("adresses", strings.Join(fpp.HostPort, ",")))

	input.Values.Set("kubeProxy.internal.clusterMasterAddresses", fpp.HostPort)

	return nil
}
