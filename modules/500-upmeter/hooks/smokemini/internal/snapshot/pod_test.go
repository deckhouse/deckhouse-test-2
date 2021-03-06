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

package snapshot

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
)

func Test_NewPodPhase(t *testing.T) {
	type args struct {
		name  string
		phase v1.PodPhase
	}
	tests := []struct {
		name string
		args args
		want PodPhase
	}{
		{
			name: "pending pod",
			args: args{name: "xx", phase: v1.PodPending},
			want: PodPhase{Name: "xx", IsPending: true},
		},
		{
			name: "not pending pod",
			args: args{name: "yy", phase: v1.PodRunning},
			want: PodPhase{Name: "yy", IsPending: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := newPodObject(tt.args.name, tt.args.phase)

			parsed, err := NewPodPhase(obj)

			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, parsed)
			}
		})
	}
}

func newPodObject(name string, phase v1.PodPhase) *unstructured.Unstructured {
	pod := v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: v1.PodStatus{
			Phase: phase,
		},
	}

	manifest, err := json.Marshal(pod)
	if err != nil {
		panic(err)
	}

	decUnstructured := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	obj := &unstructured.Unstructured{}
	if _, _, err := decUnstructured.Decode(manifest, nil, obj); err != nil {
		panic(err)
	}

	return obj
}
