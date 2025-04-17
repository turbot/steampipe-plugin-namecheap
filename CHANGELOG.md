## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#23](https://github.com/turbot/steampipe-plugin-namecheap/pull/23))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#23](https://github.com/turbot/steampipe-plugin-namecheap/pull/23))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#21](https://github.com/turbot/steampipe-plugin-namecheap/pull/21))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#21](https://github.com/turbot/steampipe-plugin-namecheap/pull/21))

## v0.2.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#17](https://github.com/turbot/steampipe-plugin-namecheap/pull/17))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#17](https://github.com/turbot/steampipe-plugin-namecheap/pull/17))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-namecheap/blob/main/docs/LICENSE). ([#17](https://github.com/turbot/steampipe-plugin-namecheap/pull/17))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#16](https://github.com/turbot/steampipe-plugin-namecheap/pull/16))

## v0.1.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#8](https://github.com/turbot/steampipe-plugin-namecheap/pull/8))

## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#6](https://github.com/turbot/steampipe-plugin-namecheap/pull/6))
- Recompiled plugin with Go version `1.21`. ([#6](https://github.com/turbot/steampipe-plugin-namecheap/pull/6))

## v0.0.2 [2023-04-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for aggregator connections not working for dynamic tables. ([#2](https://github.com/turbot/steampipe-plugin-namecheap/pull/2))

## v0.0.1 [2023-02-25]

_What's new?_

- New tables added
  - [namecheap_domain](https://hub.steampipe.io/plugins/turbot/namecheap/tables/namecheap_domain)
  - [namecheap_domain_host](https://hub.steampipe.io/plugins/turbot/namecheap/tables/namecheap_domain_host)
