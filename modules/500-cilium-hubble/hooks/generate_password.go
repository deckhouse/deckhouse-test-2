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
	"github.com/deckhouse/deckhouse/go_lib/hooks/generate_password"
)

const (
	moduleValuesKey = "ciliumHubble"
	authSecretNS    = "d8-cni-cilium"
	authSecretName  = "hubble-basic-auth"
)

var hook = generate_password.NewBasicAuthPlainHook(moduleValuesKey, authSecretNS, authSecretName)
var _ = generate_password.RegisterHook(hook)

//var _ = sdk.RegisterFunc(&go_hook.HookConfig{
//	Kubernetes: []go_hook.KubernetesConfig{
//		{
//			Name:       authSecretBinding,
//			ApiVersion: "v1",
//			Kind:       "Secret",
//			NameSelector: &types.NameSelector{
//				MatchNames: []string{authSecretName},
//			},
//			NamespaceSelector: &types.NamespaceSelector{
//				NameSelector: &types.NameSelector{
//					MatchNames: []string{authSecretNS},
//				},
//			},
//			// Synchronization is redundant because of OnBeforeHelm.
//			ExecuteHookOnSynchronization: go_hook.Bool(false),
//			ExecuteHookOnEvents:          go_hook.Bool(false),
//			FilterFunc:                   filterAuthSecret,
//		},
//	},
//	OnBeforeHelm: &go_hook.OrderedConfig{Order: 10},
//}, generatePassword)
//
//const (
//	authSecretNS      = "d8-cni-cilium"
//	authSecretName    = "hubble-basic-auth"
//	authSecretField   = "auth"
//	authSecretBinding = authSecretName
//
//	passwordValuesPath         = "ciliumHubble.auth.password"
//	passwordInternalValuesPath = "ciliumHubble.internal.auth.password"
//	externalAuthValuesPath     = "ciliumHubble.auth.externalAuthentication"
//)
//
//func filterAuthSecret(obj *unstructured.Unstructured) (go_hook.FilterResult, error) {
//	secret := &v1.Secret{}
//	err := sdk.FromUnstructured(obj, secret)
//	if err != nil {
//		return nil, fmt.Errorf("cannot convert secret to struct: %v", err)
//	}
//
//	auth := string(secret.Data[authSecretField])
//	if auth != "" {
//		auth = strings.TrimPrefix(auth, "auth:{PLAIN}")
//	}
//
//	return string(secret.Data[authSecretField]), nil
//}
//
//func generatePassword(input *go_hook.HookInput) error {
//	// Clear password from internal values if an external authentication is enabled.
//	if input.Values.Exists(externalAuthValuesPath) {
//		input.Values.Remove(passwordInternalValuesPath)
//		return nil
//	}
//
//	// Check config values.
//	password, ok := input.ConfigValues.GetOk(passwordValuesPath)
//	if ok {
//		input.Values.Set(passwordInternalValuesPath, password.String())
//		return nil
//	}
//
//	if len(input.Snapshots[authSecretBinding]) > 0 {
//		storedAuthKey := input.Snapshots[authSecretBinding][0].(string)
//		input.Values.Set(passwordInternalValuesPath, storedAuthKey)
//		return nil
//	}
//
//	// Return if auth key is already in values.
//	_, ok = input.Values.GetOk(passwordInternalValuesPath)
//	if ok {
//		return nil
//	}
//
//	// No password found, generate new one.
//	newPasswd := pwgen.AlphaNum(20)
//	input.Values.Set(passwordInternalValuesPath, newPasswd)
//	return nil
//}
