---
title: "OIDC identity provider"
permalink: en/stronghold/documentation/user/secrets-engines/identity/oidc-provider.html
---

Stronghold is an OpenID Connect ([OIDC](https://openid.net/specs/openid-connect-core-1_0.html))
identity provider. This enables client applications that speak the OIDC protocol to leverage
Stronghold's source of [identity](/docs/concepts/identity) and wide range of [authentication methods](/docs/auth)
when authenticating end-users. Client applications can configure their authentication logic
to talk to Stronghold. Once enabled, Stronghold will act as the bridge to other identity providers via
its existing authentication methods. Client applications can also obtain identity information
for their end-users by leveraging custom templating of Stronghold identity information.

{% alert level="info" %}

 **Note**: For more detailed information on the configuration resources and OIDC endpoints,
please visit the [OIDC provider](/docs/concepts/oidc-provider) concepts page.

{% endalert %}

## Setup

The Stronghold OIDC provider system is built on top of the identity secrets engine.
This secrets engine is mounted by default and cannot be disabled or moved.

Each Stronghold namespace has a default OIDC [provider](/docs/concepts/oidc-provider#providers)
and [key](/docs/concepts/oidc-provider#key). This built-in configuration enables client
applications to begin using Stronghold as a source of identity with minimal configuration. For
details on the built-in configuration and advanced options, see the [OIDC provider](/docs/concepts/oidc-provider)
concepts page.

The following steps show a minimal configuration that allows a client application to use
Stronghold as an OIDC provider.

1. Enable an Stronghold auth method:

   ```text
   $ d8 stronghold auth enable userpass
   Success! Enabled userpass auth method at: userpass/
   ```

   Any Stronghold auth method may be used within the OIDC flow. For simplicity, enable the
   `userpass` auth method.

2. Create a user:

   ```text
   $ d8 stronghold write auth/userpass/users/end-user password="securepassword"
   Success! Data written to: auth/userpass/users/end-user
   ```

   This user will authenticate to Stronghold through a client application, otherwise known as
   an OIDC [relying party](https://openid.net/specs/openid-connect-core-1_0.html#Terminology).

3. Create a client application:

   ```text
   $ d8 stronghold write identity/oidc/client/my-webapp \
     redirect_uris="https://localhost:9702/auth/oidc-callback" \
     assignments="allow_all"
   Success! Data written to: identity/oidc/client/my-webapp
   ```

   This operation creates a client application which can be used to configure an OIDC
   relying party. See the [client applications](/docs/concepts/oidc-provider#client-applications)
   section for details on different client types, including `confidential` and `public` clients.

   The `assignments` parameter limits the Stronghold entities and groups that are allowed to
   authenticate through the client application. By default, no Stronghold entities are allowed.
   To allow all Stronghold entities to authenticate, the built-in [allow_all](/docs/concepts/oidc-provider#assignments)
   assignment is provided.

4. Read client credentials:

   ```text
   $ d8 stronghold read identity/oidc/client/my-webapp

   Key                 Value
   ---                 -----
   access_token_ttl    24h
   assignments         [allow_all]
   client_id           GSDTnn3KaOrLpNlVGlYLS9TVsZgOTweO
   client_secret       hvo_secret_gBKHcTP58C4aq7FqPWsuqKgpiiegd7ahpifGae9WGkHRCwFEJTZA9KGdNVpzE0r8
   client_type         confidential
   id_token_ttl        24h
   key                 default
   redirect_uris       [https://localhost:9702/auth/oidc-callback]
   ```

   The `client_id` and `client_secret` are the client application's credentials. These
   values are typically required when configuring an OIDC relying party.

5. Read OIDC discovery configuration:

   ```text
   $ curl -s http://127.0.0.1:8200/v1/identity/oidc/provider/default/.well-known/openid-configuration
   {
     "issuer": "http://127.0.0.1:8200/v1/identity/oidc/provider/default",
     "jwks_uri": "http://127.0.0.1:8200/v1/identity/oidc/provider/default/.well-known/keys",
     "authorization_endpoint": "http://127.0.0.1:8200/ui/vault/identity/oidc/provider/default/authorize",
     "token_endpoint": "http://127.0.0.1:8200/v1/identity/oidc/provider/default/token",
     "userinfo_endpoint": "http://127.0.0.1:8200/v1/identity/oidc/provider/default/userinfo",
     "request_parameter_supported": false,
     "request_uri_parameter_supported": false,
     "id_token_signing_alg_values_supported": [
       "RS256",
       "RS384",
       "RS512",
       "ES256",
       "ES384",
       "ES512",
       "EdDSA"
     ],
     "response_types_supported": [
       "code"
     ],
     "scopes_supported": [
       "openid"
     ],
     "subject_types_supported": [
       "public"
     ],
     "grant_types_supported": [
       "authorization_code"
     ],
     "token_endpoint_auth_methods_supported": [
       "none",
       "client_secret_basic",
       "client_secret_post"
     ]
   }
   ```

   Each Stronghold OIDC provider publishes [discovery metadata](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
   The `issuer` value is typically required when configuring an OIDC relying party.

## Usage

After configuring an Stronghold auth method and client application, the following details can
be used to configure an OIDC relying party to delegate end-user authentication to Stronghold.

- `client_id` - The ID of the client application
- `client_secret` - The secret of the client application
- `issuer` - The issuer of the Stronghold OIDC provider

Otherwise, refer to the documentation of the specific OIDC relying party for usage details.

## Supported flows

The Stronghold OIDC provider feature currently supports the following authentication flow:

- [Authorization Code Flow](https://openid.net/specs/openid-connect-core-1_0.html#CodeFlowAuth).

## API

The Stronghold OIDC provider feature has a full HTTP API. Please see the
[OIDC identity provider API](/api-docs/secret/identity/oidc-provider) for more
details.
