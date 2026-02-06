# hc-auth-plugin

JWT, OAuth, and SAML authentication for the Apito API platform.

## Features

- JWT token authentication
- OAuth 2.0 integration
- SAML SSO support
- Role-based access control
- Multi-factor authentication

## Installation

```bash
go build -o hc-auth-plugin .
```

## Configuration

Configure via environment variables (see `config.yml`):

- `JWT_SECRET` - Secret for JWT signing
- `OAUTH_CLIENT_ID` - OAuth client ID
- `OAUTH_CLIENT_SECRET` - OAuth client secret

## How to Develop Plugins

See the [CDN Plugin Development Guide](https://github.com/apito-io/plugins/blob/main/CDN-PLUGIN-DEVELOPMENT.md) in the Apito plugin registry for architecture, GraphQL/REST setup, and best practices.

## How to Submit Your Plugin

1. Create your plugin repo with naming convention `hc-{name}-plugin`
2. Ensure your repo contains: `config.yml`, `README.md`, `logo.png`
3. Submit a PR to [github.com/apito-io/plugins](https://github.com/apito-io/plugins) adding your entry to `plugins.json`
