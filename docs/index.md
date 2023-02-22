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

## Documentation

- **[Table definitions & examples →](/plugins/turbot/namecheap/tables)**

## Get started

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

### Credentials from Environment Variables

The Namecheap plugin will use the standard Namecheap environment variables to obtain credentials **only if other arguments (`username`, `api_user` and `api_key`) are not specified** in the connection:

```sh
export NAMECHEAP_USERNAME=janedoe
export NAMECHEAP_API_USER=janedoe
export NAMECHEAP_API_KEY=33d0d62a6a163083ba7b3bab31bd6612
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-namecheap
- Community: [Slack Channel](https://steampipe.io/community/join)
