---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/namecheap.svg"
brand_color: "#FE5803"
display_name: "Namecheap"
short_name: "namecheap"
description: "Steampipe plugin to query domains, DNS host records and more from Namecheap."
og_description: "Query Namecheap with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/namecheap-social-graphic.png"
---

# Namecheap + Steampipe

[Namecheap](https://namecheap.com) is a domain name registrar and web hosting company

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/namecheap/tables)**

## Get started

### Install

Download and install the latest Namecheap plugin:

```bash
steampipe plugin install namecheap
```

### Configuration

Installing the latest namecheap plugin will create a config file (`~/.steampipe/config/namecheap.spc`) with a single connection named `namecheap`:

Configure your account details in `~/.steampipe/config/namecheap.spc`:

```hcl
connection "namecheap" {
  plugin = "namecheap"

  # Authentication information
  username  = "janedoe"
  api_key   = "33d0d62a6a163083ba7b3bab31bd6612"

  # IP address of the client making the request,
  # must be granted permission in Namecheap
  client_ip = "1.2.3.4"
}
```

Environment variables are also available as an alternate configuration method:
- `NAMECHEAP_USERNAME`
- `NAMECHEAP_API_USER`
- `NAMECHEAP_API_KEY`
- `NAMECHEAP_CLIENT_IP`

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-namecheap
- Community: [Slack Channel](https://steampipe.io/community/join)
