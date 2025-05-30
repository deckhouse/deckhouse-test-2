## Patches

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
