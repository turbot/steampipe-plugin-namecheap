---
title: "Steampipe Table: namecheap_domain_host - Query Namecheap Domain Hosts using SQL"
description: "Allows users to query Namecheap Domain Hosts, specifically the host records associated with a domain, providing insights into the DNS settings and configurations."
---

# Table: namecheap_domain_host - Query Namecheap Domain Hosts using SQL

Namecheap Domain Hosts is a resource within the Namecheap domain registration and management service that allows users to manage and configure host records for their domains. These host records, also known as DNS records, are essential for directing internet traffic to the correct servers and services associated with a domain. As such, they play a crucial role in domain name resolution, website hosting, email delivery, and other domain-related functions.

## Table Usage Guide

The `namecheap_domain_host` table provides insights into the DNS host records configured for domains within the Namecheap service. As a system administrator or DevOps engineer, explore host record-specific details through this table, including the associated domain, record type, address, and TTL (Time to Live) values. Utilize it to monitor and manage DNS configurations, ensure proper domain name resolution, and maintain the overall health and performance of your domains.

## Examples

### List domain host records
Discover the segments that contain specific details of a domain's host records, such as its name and address. This can be useful in managing and organizing your domains, particularly in identifying which domains have specific host settings.

```sql+postgres
select
  domain,
  name,
  type,
  address,
  ttl,
  mx_pref
from
  namecheap_domain_host
order by
  domain,
  name;
```

```sql+sqlite
select
  domain,
  name,
  type,
  address,
  ttl,
  mx_pref
from
  namecheap_domain_host
order by
  domain,
  name;
```

### List host records for a given domain
Explore the host records associated with a particular domain to gain insights into its configuration and settings. This can be especially useful for managing and troubleshooting domain-related issues.

```sql+postgres
select
  domain,
  name,
  type,
  address,
  ttl,
  mx_pref
from
  namecheap_domain_host
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  name,
  type,
  address,
  ttl,
  mx_pref
from
  namecheap_domain_host
where
  domain = 'steampipe.io';
```

### List CNAME records
Explore which domain names are associated with specific addresses, useful for identifying potential redirections or aliases. This can provide insights into your domain hosting configuration for better management and security.

```sql+postgres
select
  domain,
  name,
  address,
  ttl
from
  namecheap_domain_host
where
  type = 'CNAME';
```

```sql+sqlite
select
  domain,
  name,
  address,
  ttl
from
  namecheap_domain_host
where
  type = 'CNAME';
```