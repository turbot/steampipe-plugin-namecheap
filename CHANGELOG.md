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
