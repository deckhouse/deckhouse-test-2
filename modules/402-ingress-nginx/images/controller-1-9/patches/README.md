## Patches

## ingress-nginx

### 001-lua-info.patch

There are two patches to fix the problem with ingress names in logs and metrics.
Unfortunately, the PR was declined in the upstream.
https://github.com/kubernetes/ingress-nginx/pull/4367

### 002-makefile.patch

Run the build locally, not inside the container.

### 003-healthcheck.patch

After catching SIGTERM, ingress stops responding to the readiness probe.
The combination of this patch and the `EndpointSliceTerminatingCondition` feature gate for kube-proxy helps us avoid
traffic loss on rollout updates.
Update: for external load balancers it's advisable to get 5xx if a SIGTERM was sent to the controller, we control this logic by applying/checking `D8s-External-Check` http header.

Backport of the behavior of the later versions of ingress nginx controller.
The `sleep` is needed to gracefully shut down ingress controllers behind a cloud load balancer.

### 004-metrics-SetSSLExpireTime.patch

Fixes namespace which is given by metric nginx_ingress_controller_ssl_expire_time_seconds.

https://github.com/kubernetes/ingress-nginx/pull/10274

### 005-util.patch

Adds "-e /dev/null" flags to the "nginx -t" invocations so that "nginx -t" logs aren't got saved to /var/log/nginx/error.log file, preventing fs bloating.

### 006-fix-cleanup.patch

Fix tmpDir path for the cleanup procedure.

https://github.com/kubernetes/ingress-nginx/pull/10797

### 007-geoip.patch

https://github.com/kubernetes/ingress-nginx/pull/10495

### 008-new-metrics.patch

This patch adds worker max connections, worker processes and worker max open files metrics.

### 009-default-backend-fix.patch

Fixes the problem with the controller when Ingress specifies `Service` with the `ExternalName` type as the main backend, and the default backend (using an annotation `nginx.ingress.kubernetes.io/default-backend `) - with the ClusterIP type. You can see the detailed cases here:
https://github.com/kubernetes/ingress-nginx/issues/12158
https://github.com/kubernetes/ingress-nginx/issues/12173
https://github.com/deckhouse/deckhouse/issues/9933

### 010-cve.patch

Fix cve vulnerabilities.

### 011-nginx-build.patch

Build nginx for controller on ALT Linux.

## rootfs

### 001-balancer-lua.patch

TODO: update readme with patch description

### 002-nginx-tmpl.patch

* Enable our metrics collector instead of the default one.
* Enable pcre_jit.
* Add the health check server to provide the way for an external load balancer to check that the ingress controller will be terminated soon.
* Set default values for upstream_retries and total_upstream_response_time to avoid incorrect logs when it is a raw tcp request.
* Set proxy connect timeout for auth locations.
* Replace the status field with formatted status field which is explicitly converted to number to avoid incorrect logs when response status is 009.

We do not intend to make a PR to the upstream with these changes, because there are only our custom features.

### 003-auth-cookie-always.patch

Without always option toggled, ingress-nginx does not set the cookie in case if backend returns >=400 code, which may lead to dex refresh token invalidation.
Annotation `nginx.ingress.kubernetes.io/auth-always-set-cookie` does not work. Anyway, we can't use it, because we need this behavior for all ingresses.

https://github.com/kubernetes/ingress-nginx/pull/8213
