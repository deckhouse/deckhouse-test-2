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
	"context"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/cloudflare/cfssl/csr"
	"github.com/flant/addon-operator/pkg/module_manager/go_hook"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"github.com/deckhouse/deckhouse/go_lib/certificate"
	. "github.com/deckhouse/deckhouse/testing/hooks"
)

var _ = Describe("Cert Manager hooks :: generate_webhook_certs ::", func() {
	f := HookExecutionConfigInit(`{"certManager":{"internal":{}}}`, "")

	Context("Without secret", func() {
		BeforeEach(func() {
			f.BindingContexts.Set(f.GenerateBeforeHelmContext())
			f.RunHook()
		})

		It("Should add ca and certificate to values", func() {
			Expect(f).To(ExecuteSuccessfully())
			Expect(f.BindingContexts.Array()).ShouldNot(BeEmpty())

			Expect(f.ValuesGet("certManager.internal.webhookCACrt").Exists()).To(BeTrue())
			Expect(f.ValuesGet("certManager.internal.webhookCAKey").Exists()).To(BeTrue())
			Expect(f.ValuesGet("certManager.internal.webhookCrt").Exists()).To(BeTrue())
			Expect(f.ValuesGet("certManager.internal.webhookKey").Exists()).To(BeTrue())

			blockCA, _ := pem.Decode([]byte(f.ValuesGet("certManager.internal.webhookCACrt").String()))
			certCA, err := x509.ParseCertificate(blockCA.Bytes)
			Expect(err).To(BeNil())
			Expect(certCA.IsCA).To(BeTrue())
			Expect(certCA.Subject.CommonName).To(Equal("cert-manager-webhook"))

			block, _ := pem.Decode([]byte(f.ValuesGet("certManager.internal.webhookCrt").String()))
			cert, err := x509.ParseCertificate(block.Bytes)
			Expect(err).To(BeNil())
			Expect(cert.IsCA).To(BeFalse())
			Expect(cert.Subject.CommonName).To(Equal("cert-manager-webhook"))
		})
	})
	Context("With secrets", func() {
		caAuthority, _ := genWebhookCa(nil)
		tlsAuthority, _ := genWebhookTLS(&go_hook.HookInput{LogEntry: logrus.New().WithContext(context.Background())}, caAuthority)

		BeforeEach(func() {
			f.BindingContexts.Set(f.KubeStateSet(fmt.Sprintf(`
---
apiVersion: v1
kind: Secret
type: kubernetes.io/tls
metadata:
  name: cert-manager-webhook-ca
  namespace: d8-cert-manager
data:
  ca.crt: %[1]s
  tls.crt: %[2]s
  tls.key: %[3]s
---
apiVersion: v1
kind: Secret
type: kubernetes.io/tls
metadata:
  name: cert-manager-webhook-tls
  namespace: d8-cert-manager
data:
  ca.crt: %[1]s
  tls.crt: %[2]s
  tls.key: %[3]s
`, base64.StdEncoding.EncodeToString([]byte(caAuthority.Cert)),
				base64.StdEncoding.EncodeToString([]byte(tlsAuthority.Cert)),
				base64.StdEncoding.EncodeToString([]byte(tlsAuthority.Key)))),
			)
			f.RunHook()
		})
		It("Should add existing ca certificate to values", func() {
			Expect(f).To(ExecuteSuccessfully())
			Expect(f.BindingContexts.Array()).ShouldNot(BeEmpty())

			Expect(f.ValuesGet("certManager.internal.webhookCACrt").String()).To(Equal(caAuthority.Cert))
			Expect(f.ValuesGet("certManager.internal.webhookCAKey").String()).To(Equal(tlsAuthority.Key))
			Expect(f.ValuesGet("certManager.internal.webhookCrt").String()).To(Equal(tlsAuthority.Cert))
			Expect(f.ValuesGet("certManager.internal.webhookKey").String()).To(Equal(tlsAuthority.Key))
		})
	})

	Context("With legacy secrets", func() {
		caAuthority := testGenerateLegacy()
		tlsAuthority, _ := genWebhookTLS(&go_hook.HookInput{
			LogEntry: logrus.New().WithContext(context.Background()),
		}, &caAuthority)

		BeforeEach(func() {
			f.BindingContexts.Set(f.KubeStateSet(fmt.Sprintf(`
---
apiVersion: v1
kind: Secret
type: kubernetes.io/tls
metadata:
  name: cert-manager-webhook-ca
  namespace: d8-cert-manager
data:
  ca.crt: %[1]s
  tls.crt: %[2]s
  tls.key: %[3]s
---
apiVersion: v1
kind: Secret
type: kubernetes.io/tls
metadata:
  name: cert-manager-webhook-tls
  namespace: d8-cert-manager
data:
  ca.crt: %[1]s
  tls.crt: %[2]s
  tls.key: %[3]s
`,
				base64.StdEncoding.EncodeToString([]byte(caAuthority.Cert)),
				base64.StdEncoding.EncodeToString([]byte(tlsAuthority.Cert)),
				base64.StdEncoding.EncodeToString([]byte(tlsAuthority.Key)))),
			)
			f.RunHook()
		})
		It("Should generate new ca and put to values", func() {
			Expect(f).To(ExecuteSuccessfully())
			Expect(f.BindingContexts.Array()).ShouldNot(BeEmpty())

			Expect(f.ValuesGet("certManager.internal.webhookCACrt").Exists()).To(BeTrue())
			Expect(f.ValuesGet("certManager.internal.webhookCAKey").Exists()).To(BeTrue())
			Expect(f.ValuesGet("certManager.internal.webhookCrt").Exists()).To(BeTrue())
			Expect(f.ValuesGet("certManager.internal.webhookKey").Exists()).To(BeTrue())

			Expect(f.ValuesGet("certManager.internal.webhookCACrt").String()).ToNot(Equal(caAuthority.Cert))
			Expect(f.ValuesGet("certManager.internal.webhookCAKey").String()).ToNot(Equal(tlsAuthority.Key))
			Expect(f.ValuesGet("certManager.internal.webhookCrt").String()).ToNot(Equal(tlsAuthority.Cert))
			Expect(f.ValuesGet("certManager.internal.webhookKey").String()).ToNot(Equal(tlsAuthority.Key))
		})
	})
})

func testGenerateLegacy() certificate.Authority {
	ca, _ := certificate.GenerateCA(nil, "cert-manager.webhook.ca", func(r *csr.CertificateRequest) {
		r.KeyRequest = &csr.KeyRequest{
			A: "rsa",
			S: 2048,
		}
		r.Hosts = []string{
			"cert-manager-webhook.d8-cert-manager.svc",
		}
		r.Names = []csr.Name{
			{O: "cert-manager-webhook.d8-cert-manager"},
		}
	})

	return ca
}
