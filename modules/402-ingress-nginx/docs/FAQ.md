---
title: "The ingress-nginx module: FAQ"
---

## How do I limit access to the application in the cluster to Ingress controllers only?

If you need to limit access to your application inside the cluster exclusively from Ingress Pods, you should add a kube-rbac-proxy container to the application Pod as shown in the example below:

{% raw %}

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
  namespace: my-namespace
spec:
  selector:
    matchLabels:
      app: my-app
  replicas: 1
  template:
    metadata:
      labels:
        app: my-app
    spec:
      serviceAccountName: my-sa
      containers:
      - name: my-cool-app
        image: mycompany/my-app:v0.5.3
        args:
        - "--listen=127.0.0.1:8080"
        livenessProbe:
          httpGet:
            path: /healthz
            port: 443
            scheme: HTTPS
      - name: kube-rbac-proxy
        image: flant/kube-rbac-proxy:v0.1.0
        # It is recommended to use a proxy from Deckhouse repository.
        args:
        - "--secure-listen-address=0.0.0.0:443"
        - "--config-file=/etc/kube-rbac-proxy/config-file.yaml"
        - "--v=2"
        - "--logtostderr=true"
        # The user authentication and authorization are not possible if the kube-apiserver is not available.
        # Stale Cache stores the results of successful authorization and is used only if the apiserver is not available.
        - "--stale-cache-interval=1h30m"
        ports:
        - containerPort: 443
          name: https
        volumeMounts:
        - name: kube-rbac-proxy
          mountPath: /etc/kube-rbac-proxy
      volumes:
      - name: kube-rbac-proxy
        configMap:
          name: kube-rbac-proxy
```

{% endraw %}

The application only accepts localhost (127.0.0.1) requests. That means that an unsecured connection can only be established to it from within the Pod.
At the same time, the proxy listens on 0.0.0.0 and intercepts all external traffic to the Pod.

### How do I provide minimum rights to the Service Account?

The proxy needs permissions to create `TokenReview` and `SubjectAccessReview` to authenticate and authorize users using the kube-apiserver.

DKP clusters [already have a ready-made ClusterRole](https://github.com/deckhouse/deckhouse/blob/main/modules/002-deckhouse/templates/common/rbac/kube-rbac-proxy.yaml) — **d8-rbac-proxy**, you don't need to create it yourself!
Link it to your Deployment's ServiceAccount, as shown in the example below.
{% raw %}

```yaml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-sa
  namespace: my-namespace
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: my-namespace:my-sa:d8-rbac-proxy
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: d8:rbac-proxy
subjects:
- kind: ServiceAccount
  name: my-sa
  namespace: my-namespace
```

### The Kube-RBAC-Proxy configuration

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-rbac-proxy
data:
  config-file.yaml: |+
    excludePaths:
    - /healthz
  # No authorization for liveness probes is required.
    upstreams:
    - upstream: http://127.0.0.1:8081/
  # The destination address
      path: /
  # The path to the proxy to forward requests to the upstream.
      authorization:
        resourceAttributes:
          namespace: my-namespace
          apiGroup: apps
          apiVersion: v1
          resource: deployments
          subresource: http
          name: my-app
```

{% endraw %}
According to the configuration, the user must have access to the `my-app` Deployment and its `http` subresource in the `my-namespace` namespace.

Such permissions have the following RBAC form:
{% raw %}

```yaml
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: kube-rbac-proxy:my-app
  namespace: my-namespace
rules:
- apiGroups: ["apps"]
  resources: ["deployments/http"]
  resourceNames: ["my-app"]
  verbs: ["get", "create", "update", "patch", "delete"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: kube-rbac-proxy:my-app
  namespace: my-namespace
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: kube-rbac-proxy:my-app
subjects:
# All user certificates of Ingress controllers are issued for one specific group.
- kind: Group
  name: ingress-nginx:auth
```

For the Ingress resource, add parameters:

```yaml
nginx.ingress.kubernetes.io/backend-protocol: HTTPS
nginx.ingress.kubernetes.io/configuration-snippet: |
  proxy_ssl_certificate /etc/nginx/ssl/client.crt;
  proxy_ssl_certificate_key /etc/nginx/ssl/client.key;
  proxy_ssl_protocols TLSv1.2;
  proxy_ssl_session_reuse on;
```

{% endraw %}
[Here](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#x509-client-certs) you can read more about how certificate authentication works.

## How do I configure an external load balancer to check if IngressNginxController is available?

In case an `IngressNginxController` is deployed behind a load balancer, it is advisable to configure your load balancer so that it would check
the availability of the IngressNginxController's endpoints via a health check mechanism, periodically sending either HTTP-requests or TCP-packets.
While it is possible to test the endpoints simply by checking if a relevant TCP port is open, we recommend implementing HTTP checks with the following parameters:

- Protocol: `HTTP`
- Path: `/healthz`
- Port: `80` (or relevant [httpPort](cr.html#ingressnginxcontroller-v1-spec-hostport-httpport) value in case of using `HostPort` inlet).

## How do I configure MetalLB to be accessible from the internal network only?

Below is an example of a MetalLB config with access from the internal network only:

```yaml
apiVersion: deckhouse.io/v1
kind: IngressNginxController
metadata:
  name: main
spec:
  ingressClass: "nginx"
  inlet: "LoadBalancer"
  loadBalancer:
    sourceRanges:
    - 192.168.0.0/24
```

{% alert level="warning" %}
The [`svcSourceRangeCheck`](../cni-cilium/configuration.html#parameters-svcsourcerangecheck) parameter should be enabled in cni-cilium module for correct work.
{% endalert %}

## How to add extra log fields to a nginx-controller?

Example of adding extra fields:

```yaml
apiVersion: deckhouse.io/v1
kind: IngressNginxController
metadata:
  name: main
spec:
  ingressClass: "nginx"
  inlet: "LoadBalancer"
  additionalLogFields:
    my-cookie: "$cookie_MY_COOKIE"
```

## How to enable HorizontalPodAutoscaling for IngressNginxController?

{% alert level="warning" %}
HPA mode is possible only for controllers with inlet: `LoadBalancer` or `LoadBalancerWithProxyProtocol`.

HPA mode is possible only for `minReplicas` != `maxReplicas` otherwise deployment `hpa-scaler` will not be created.
{% endalert %}

HPA is set with attributes `minReplicas` and `maxReplicas` in a [IngressNginxController CR](cr.html#ingressnginxcontroller).

The IngressNginxController is deployed using DaemonSet. DaemonSet does not provide horizontal scaling capabilities, so `hpa-scaler` Deployment will be created with the HPA resource, which is observing custom metric `prometheus-metrics-adapter-d8-ingress-nginx-cpu-utilization-for-hpa`. If CPU utilization exceeds 50%, the HPA-controller scales `hpa-scaler` Deployment with a new replica (with respect to `minReplicas` and `maxReplicas`).

`hpa-scaler` Deployment has HardPodAntiAffinity, and it will order a new Node (inside its NodeGroup), where one more ingress-controller will be set.

Notes:

- The minimum actual number of IngressNginxController replicas cannot be less than the minimum number of nodes in the node group it is deployed to.
- The maximum actual number of IngressNginxController replicas cannot be greater than the maximum number of nodes in the node group it is deployed to.

## How to use IngressClass with IngressClassParameters?

Since version 1.1 IngressNginxController Deckhouse creates an IngressClass object. If you want to use your own IngressClass
with your customized IngressClassParameters, you need to add the label `ingress-class.deckhouse.io/external: "true"`:

```yaml
apiVersion: networking.k8s.io/v1
kind: IngressClass
metadata:
  labels:
    ingress-class.deckhouse.io/external: "true"
  name: my-super-ingress
spec:
  controller: ingress-nginx.deckhouse.io/my-super-ingress
  parameters:
    apiGroup: elbv2.k8s.aws
    kind: IngressClassParams
    name: awesome-class-cfg
```

In this case Deckhouse will not create an IngressClass object and will use your own.

## How to disable the collection of detailed Ingress resources statistics?

By default, Deckhouse collects detailed statistics from all Ingress resources in the cluster, which generates a high load on the monitoring system.

To disable statistics collection, add label `ingress.deckhouse.io/discard-metrics: "true"` to the corresponding Namespace or Ingress resource.

Example of disabling statistics (metrics) collection for all Ingress resources in the `review-1` namespace:

```shell
kubectl label ns review-1 ingress.deckhouse.io/discard-metrics=true
```

Example of disabling statistics (metrics) collection for all `test-site` Ingress resources in the `development` namespace:

```shell
kubectl label ingress test-site -n development ingress.deckhouse.io/discard-metrics=true
```

## How do I correctly drain a node running an IngressNginxController's pods?

There are two ways to gracefully drain a node running IngressNginxController.

1. Using an annotation.

    The annotation will be automatically removed after the operation completes.

    ```shell
    kubectl annotate node <node_name> update.node.deckhouse.io/draining=user
    ```

1. Using kubectl drain.

    When using the standard kubectl drain command, you must specify the `--force` flag even if `--ignore-daemonsets` is present,
    since IngressNginxController is deployed using Advanced DaemonSet:

    ```shell
    kubectl drain <node_name> --delete-emptydir-data --ignore-daemonsets --force
    ```

## How to enable Web Application Firewall (WAF)?

Software known as a Web Application Firewall (WAF) is used to protect web applications from Layer 7 attacks.
Ingress-nginx controller has a built-in WAF called `ModSecurity' (The Open Worldwide Application Security Project).

ModSecurity is disabled by default.

### Enabling ModSecurity

To enable ModSecurity, you must specify the following parameters in the `config` section of the CR IngressNginxController:

```yaml
apiVersion: deckhouse.io/v1
kind: IngressNginxController
metadata:
  name: <name_of_the_controller>
spec:
  config:
    enable-modsecurity: "true"
    modsecurity-snippet: |
      Include /etc/nginx/modsecurity/modsecurity.conf
```

After applying the settings, ModSecurity will start working for all traffic passing through this Ingress nginx controller.
This uses the audit mode (`DetectionOnly`) and [basic recommended configuration](https://github.com/owasp-modsecurity/ModSecurity/blob/v3/master/modsecurity.conf-recommended).

### Setting up ModSecurity

You can configure ModSecurity in two ways:
1. For the entire ingress-nginx controller
   - the necessary directives are described in the section `config.modsecurity-snippet` in CR IngressNginxController, as in the example above.
1. For each CR Ingress separately
   - the necessary directives are described in the annotation `nginx.ingress.kubernetes.io/modsecurity-snippet : |` directly in Ingress manifests.

To enable the execution of rules (and not just logging), add the `SecRuleEngine On` directive as shown below:

```yaml
apiVersion: deckhouse.io/v1
kind: IngressNginxController
metadata:
  name: <name_of_the_controller>
spec:
  config:
    enable-modsecurity: "true"
    modsecurity-snippet: |
      Include /etc/nginx/modsecurity/modsecurity.conf
      SecRuleEngine On
```

A full list and description of the directives can be found in the [official documentation](https://github.com/owasp-modsecurity/ModSecurity/wiki/Reference-Manual-%28v3.x%29 ).

Currently, the OWASP Core Rule Set (CRS) is not available.
