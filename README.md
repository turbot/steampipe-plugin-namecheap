![image](https://hub.steampipe.io/images/plugins/turbot/namecheap-social-graphic.png)

# Namecheap Plugin for Steampipe

Use SQL to query models, completions and more from Namecheap.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/namecheap)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/namecheap/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-namecheap/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install namecheap
```

Configure your account details in `~/.steampipe/config/namecheap.spc`:

```hcl
connection "namecheap" {
  # Authentication information
  username  = "janedoe"
  api_key   = "33d0d62a6a163083ba7b3bab31bd6612"
  # IP address of the client making the request,
  # must be granted permission in Namecheap
  client_ip = "1.2.3.4"
}
```

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
  domain
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
