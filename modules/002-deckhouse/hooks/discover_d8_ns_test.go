/*
Copyright 2023 Flant JSC

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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/deckhouse/deckhouse/testing/hooks"
)

var _ = Describe("Modules :: flow-schema :: hooks :: discover deckhouse namespaces ::", func() {
	f := HookExecutionConfigInit(`{"deckhouse":{"internal": {"namespaces": []}}}`, `{}`)

	Context("Empty cluster", func() {
		BeforeEach(func() {
			f.KubeStateSet(``)
			f.BindingContexts.Set(f.GenerateBeforeHelmContext())
			f.RunHook()
		})

		It("Hook should execute successfully", func() {
			Expect(f).To(ExecuteSuccessfully())
			Expect(f.ValuesGet(namespacesValuesPath).String()).To(MatchJSON(`[]`))
		})
	})

	Context("Cluster with namespaces", func() {
		BeforeEach(func() {
			f.KubeStateSet(`
---
apiVersion: v1
kind: Namespace
metadata:
  name: test2
  labels:
    heritage: deckhouse
---
apiVersion: v1
kind: Namespace
metadata:
  name: test1
  labels:
    heritage: deckhouse
---
apiVersion: v1
kind: Namespace
metadata:
  name: test3
---
apiVersion: v1
kind: Namespace
metadata:
  name: test4
  labels:
    heritage: upmeter
`)
			f.BindingContexts.Set(f.GenerateBeforeHelmContext())
			f.RunHook()
		})

		It("Values must be set", func() {
			Expect(f).To(ExecuteSuccessfully())
			Expect(f.ValuesGet(namespacesValuesPath).String()).To(MatchJSON(`["test1","test2"]`))
		})
	})
})
