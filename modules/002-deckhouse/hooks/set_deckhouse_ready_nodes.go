/*
Copyright 2024 Flant JSC

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
	"fmt"
	"log/slog"
	"strconv"

	"github.com/flant/addon-operator/pkg/module_manager/go_hook"
	"github.com/flant/addon-operator/sdk"
	"github.com/flant/shell-operator/pkg/kube_events_manager/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/kubectl/pkg/util/podutils"
	"k8s.io/utils/ptr"

	sdkobjectpatch "github.com/deckhouse/module-sdk/pkg/object-patch"
)

const deckhouseReadyLabel = "node.deckhouse.io/deckhouse-ready"

var _ = sdk.RegisterFunc(&go_hook.HookConfig{
	Settings:     &go_hook.HookConfigSettings{},
	Queue:        "/modules/deckhouse/set-deckhouse-ready-nodes",
	OnBeforeHelm: &go_hook.OrderedConfig{Order: 1},
	Kubernetes: []go_hook.KubernetesConfig{
		{
			Name:                         "control-plane-pods",
			ApiVersion:                   "v1",
			Kind:                         "Pod",
			ExecuteHookOnSynchronization: ptr.To(true),
			ExecuteHookOnEvents:          ptr.To(true),
			NamespaceSelector: &types.NamespaceSelector{
				NameSelector: &types.NameSelector{
					MatchNames: []string{"kube-system"},
				},
			},
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"tier":      "control-plane",
					"component": "kube-apiserver",
				},
			},
			FilterFunc: applyPodFilter,
		},
		{
			Name:                         "control-plane-nodes",
			ApiVersion:                   "v1",
			Kind:                         "Node",
			ExecuteHookOnSynchronization: ptr.To(true),
			ExecuteHookOnEvents:          ptr.To(true),
			LabelSelector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"node-role.kubernetes.io/control-plane": "",
				},
			},
			FilterFunc: applyNodeFilter,
		},
	},
}, setDeckhouseReadyNodes)

type statusPod struct {
	Name    string
	Node    string
	IsReady bool
}

type statusNode struct {
	Name    string
	IsReady bool
}

func applyNodeFilter(obj *unstructured.Unstructured) (go_hook.FilterResult, error) {
	var node corev1.Node

	err := sdk.FromUnstructured(obj, &node)
	if err != nil {
		return nil, fmt.Errorf("cannot convert kubernetes object: %v", err)
	}

	var isReady bool
	for _, cond := range node.Status.Conditions {
		if cond.Type == corev1.NodeReady && cond.Status == corev1.ConditionTrue {
			isReady = true
			break
		}
	}

	return statusNode{
		Name:    node.Name,
		IsReady: isReady,
	}, nil
}

func applyPodFilter(obj *unstructured.Unstructured) (go_hook.FilterResult, error) {
	var pod corev1.Pod

	err := sdk.FromUnstructured(obj, &pod)
	if err != nil {
		return nil, fmt.Errorf("cannot convert kubernetes object: %v", err)
	}

	return statusPod{
		Name:    pod.Name,
		Node:    pod.Spec.NodeName,
		IsReady: podutils.IsPodReady(&pod),
	}, nil
}

func setDeckhouseReadyNodes(input *go_hook.HookInput) error {
	pods, err := sdkobjectpatch.UnmarshalToStruct[statusPod](input.NewSnapshots, "control-plane-pods")
	if err != nil {
		return fmt.Errorf("failed to unmarshal control-plane-pods snapshot: %w", err)
	}

	nodes, err := sdkobjectpatch.UnmarshalToStruct[statusNode](input.NewSnapshots, "control-plane-nodes")
	if err != nil {
		return fmt.Errorf("failed to unmarshal control-plane-nodes snapshot: %w", err)
	}

	if len(nodes) == 0 {
		return nil
	}

	podPerNode := make(map[string]bool, len(pods))
	for _, p := range pods {
		podPerNode[p.Node] = p.IsReady
	}

	deckhouseReadyNodes := make(map[string]bool, 0)
	for _, n := range nodes {
		if !n.IsReady {
			deckhouseReadyNodes[n.Name] = false
			continue
		}

		deckhouseReadyNodes[n.Name] = podPerNode[n.Name]
	}

	for nodeName, nodeStatus := range deckhouseReadyNodes {
		input.Logger.Info("Labeling node with label", slog.String("label_key", nodeName), slog.String("label_value", deckhouseReadyLabel), slog.Bool("status", nodeStatus))
		metadata := map[string]interface{}{
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					deckhouseReadyLabel: strconv.FormatBool(nodeStatus),
				},
			},
		}
		input.PatchCollector.PatchWithMerge(metadata, "v1", "Node", "", nodeName)
	}

	return nil
}
