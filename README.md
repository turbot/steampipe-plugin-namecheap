![image](https://hub.steampipe.io/images/plugins/turbot/namecheap-social-graphic.png)

# Namecheap Plugin for Steampipe

Use SQL to query models, completions and more from Namecheap.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/namecheap)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/namecheap/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-namecheap/issues)

## Quick start

### Install

Download and install the latest Namecheap plugin:

```bash
steampipe plugin install namecheap
```

### Credentials

| Item        | Description                                                                                                                                                                                           |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Namecheap requires a `username` or `API user` and an [API key](https://www.namecheap.com/support/api/intro/) for all requests.                                                                |
| Permissions | API keys have the same permissions as the user who creates them, and if the user permissions change, the API key permissions also change.                                                         |
| Radius      | Each connection represents a single Namecheap Installation.                                                                                                                                           |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/namecheap.spc`)<br />2. Credentials specified in environment variables, e.g., `NAMECHEAP_USERNAME`, `NAMECHEAP_API_USER` and `NAMECHEAP_API_KEY`. |

### Configuration

Installing the latest namecheap plugin will create a config file (`~/.steampipe/config/namecheap.spc`) with a single connection named `namecheap`:

Configure your account details in `~/.steampipe/config/namecheap.spc`:

```hcl
connection "namecheap" {
  plugin = "namecheap"

  # Username/API User is required for requests. Required.
  # This can also be set via the `NAMECHEAP_USERNAME` environment variable.
  # username = "janedoe"

  # A specific API User can also be defined. Optional, by default, this will be
  # set to the `username`.
  # This can also be set via the `NAMECHEAP_API_USER` environment variable.
  # api_user = "janedoe"

  # API key for requests. Required.
  # See instructions at https://www.namecheap.com/support/api/intro/.
  # This can also be set via the `NAMECHEAP_API_KEY` environment variable.
  # api_key = "33d0d62a6a163083ba7b3bab31bd6612"
}
```

Alternatively, you can also use the standard Namecheap environment variables to obtain credentials **only if other arguments (`username`, `api_user` and `api_key`) are not specified** in the connection:

```sh
export NAMECHEAP_USERNAME=janedoe
export NAMECHEAP_API_USER=janedoe
export NAMECHEAP_API_KEY=33d0d62a6a163083ba7b3bab31bd6612
```

### Execution

Run steampipe:

```shell
steampipe query
```

List domains in your Namecheap account:

```sql
select
  domain,
  expires
from
  namecheap_domain
order by
  domain;
```

```
+--------------+---------------------------+
| domain       | expires                   |
+--------------+---------------------------+
| steampipe.io | 2025-04-19T20:00:00-04:00 |
| turbot.com   | 2024-03-07T20:00:00-04:00 |
+--------------+---------------------------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-namecheap.git
cd steampipe-plugin-namecheap
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/namecheap.spc
```

Try it!

```
steampipe query
> .inspect namecheap
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-namecheap/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Namecheap Plugin](https://github.com/turbot/steampipe-plugin-namecheap/labels/help%20wanted)
