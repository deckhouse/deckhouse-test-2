/*
Copyright 2022 Flant JSC

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

	"github.com/deckhouse/deckhouse/go_lib/deckhouse-config/modules"
	. "github.com/deckhouse/deckhouse/testing/hooks"
)

var _ = Describe("DeckhouseConfig hooks :: discover", func() {
	f := HookExecutionConfigInit(`"deckhouseConfig": {"internal":{}}`, `{}`)

	registryErr := modules.Registry().Init(modules.DefaultGlobalHooksDir, modules.DefaultModulesDir)

	Context("Empty cluster", func() {
		BeforeEach(func() {
			f.KubeStateSet(``)
			f.BindingContexts.Set(f.GenerateOnStartupContext())
			f.RunHook()
		})

		It("All module names should be in values", func() {
			Expect(registryErr).ShouldNot(HaveOccurred(), "modules registry should be inited")
			Expect(f).To(ExecuteSuccessfully())

			Expect(f.ValuesGet(PossibleNamesPath).Exists()).To(BeTrue())
		})
	})
})
