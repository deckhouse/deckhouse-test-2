/*
Copyright 2023 Flant JSC
Licensed under the Deckhouse Platform Enterprise Edition (EE) license. See https://github.com/deckhouse/deckhouse/blob/main/ee/LICENSE
*/

package hooks

import (
	"crypto/md5"
	"fmt"

	"github.com/flant/addon-operator/pkg/module_manager/go_hook"
	"github.com/flant/addon-operator/sdk"
	"github.com/flant/shell-operator/pkg/kube_events_manager/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/deckhouse/deckhouse/pkg/log"
)

const (
	ns     = "d8-user-authn"
	cmName = "whitelabel-custom-logo"
)

var _ = sdk.RegisterFunc(&go_hook.HookConfig{
	Queue:        "/modules/user-authn/custom_logo",
	OnBeforeHelm: &go_hook.OrderedConfig{Order: 10},
	Kubernetes: []go_hook.KubernetesConfig{
		{
			Name:       "logo-cm",
			ApiVersion: "v1",
			Kind:       "ConfigMap",
			FilterFunc: filterLogoCM,
			NamespaceSelector: &types.NamespaceSelector{
				NameSelector: &types.NameSelector{
					MatchNames: []string{"d8-system"},
				},
			},
			NameSelector: &types.NameSelector{
				MatchNames: []string{cmName},
			},
		},
	},
}, customLogoHandler)

func customLogoHandler(input *go_hook.HookInput) error {
	if !input.Values.Get("global.clusterIsBootstrapped").Bool() {
		input.Logger.Info("Cluster is not yet bootstrapped, skipping custom logo")
		return nil
	}

	snaps := input.NewSnapshots.Get("logo-cm")
	if len(snaps) == 0 {
		input.Values.Set("userAuthn.internal.customLogo.enabled", false)
		input.PatchCollector.DeleteInBackground("v1", "ConfigMap", ns, cmName)
		return nil
	}

	logoData := new(logos)

	err := snaps[0].UnmarshalTo(logoData)
	if err != nil {
		input.Logger.Warn("cannot unmarshal logo data", log.Err(err))
		return nil
	}

	cm := buildDexLogoCM(logoData.Logo, logoData.Title)

	sum := logoData.Logo + logoData.Title

	md5Sum := md5.Sum([]byte(sum))

	input.PatchCollector.CreateOrUpdate(cm)
	input.Values.Set("userAuthn.internal.customLogo.enabled", true)
	input.Values.Set("userAuthn.internal.customLogo.checksum", fmt.Sprintf("%x", md5Sum))

	return nil
}

func filterLogoCM(obj *unstructured.Unstructured) (go_hook.FilterResult, error) {
	var cm v1.ConfigMap

	err := sdk.FromUnstructured(obj, &cm)
	if err != nil {
		return nil, err
	}

	logo, ok := cm.Data["dexLogo"]
	if !ok {
		return nil, nil
	}
	title, ok := cm.Data["dexTitle"]
	if !ok {
		return nil, nil
	}

	return logos{
		Logo:  logo,
		Title: title,
	}, nil
}

type logos struct {
	Logo  string
	Title string
}

func buildDexLogoCM(logo, title string) *v1.ConfigMap {
	return &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cmName,
			Namespace: ns,
			Labels: map[string]string{
				"area": "whitelabel",
			},
		},
		Data: map[string]string{
			"dexLogo":  logo,
			"dexTitle": title,
		},
	}
}
