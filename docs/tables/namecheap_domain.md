---
title: "Steampipe Table: namecheap_domain - Query Namecheap Domains using SQL"
description: "Allows users to query Namecheap Domains, specifically providing detailed information about each registered domain such as domain ID, domain name, expiration date, and more."
---

# Table: namecheap_domain - Query Namecheap Domains using SQL

Namecheap is a domain registrar and technology company that provides services for buying, hosting, and managing domains. It offers a wide range of domain management services including domain name registration, domain name transfers, domain name renewal, domain expiration protection, and domain privacy services. Namecheap's domain management platform allows users to manage their domains in a centralized manner, providing tools for DNS management, domain forwarding, email hosting, and more.

## Table Usage Guide

The `namecheap_domain` table provides insights into domains registered with Namecheap. As a system administrator, you can explore domain-specific details through this table, including domain IDs, domain names, expiration dates, and more. Utilize it to manage and monitor your domains, such as checking the expiration date of domains, verifying the status of domain transfers, and keeping track of all registered domains.

## Examples

### List domains
Explore which domains are nearing their expiration dates to ensure timely renewals and uninterrupted service.

```sql+postgres
select
  domain,
  expires
from
  namecheap_domain;
```

```sql+sqlite
select
  domain,
  expires
from
  namecheap_domain;
```

### Get a specific domain
Discover the expiration date of a specific domain to stay updated and renew it before it lapses. This query is beneficial for domain management, ensuring uninterrupted online presence.

```sql+postgres
select
  domain,
  expires
from
  namecheap_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  expires
from
  namecheap_domain
where
  domain = 'steampipe.io';
```

### Domains expiring in the next 90 days
Explore which domains are due to expire in the next 90 days. This can be particularly useful for domain renewal management, ensuring that you don't lose control of important domains due to overlooked expiration dates.

```sql+postgres
select
  domain,
  expires,
  age(expires)
from
  namecheap_domain
where
  expires < (current_date + interval '90 days');
```

```sql+sqlite
select
  domain,
  expires,
  julianday('now') - julianday(expires)
from
  namecheap_domain
where
  julianday(expires) < julianday(date('now','+90 day'));
```

### Domains by date created at Namecheap
Discover the segments that have been recently established on Namecheap. This can help users identify new domains and understand their age, facilitating better management and oversight of their online properties.

```sql+postgres
select
  domain,
  created,
  age(created)
from
  namecheap_domain
order by
  created;
```

```sql+sqlite
select
  domain,
  created,
  julianday('now') - julianday(created) as age_created
from
  namecheap_domain
order by
  created;
```

### Domains without auto-renew enabled
Discover the domains that do not have the auto-renew feature enabled. This can help manage domain expiry and avoid losing access to your domains due to missed renewals.

```sql+postgres
select
  domain,
  expires
from
  namecheap_domain
where
  not auto_renew;
```

```sql+sqlite
select
  domain,
  expires
from
  namecheap_domain
where
  auto_renew = 0;
```

### Domains and their nameservers
Explore which domains are associated with specific nameservers. This is useful in managing and understanding the distribution of your domains across different nameservers.

```sql+postgres
select
  domain,
  ns
from
  namecheap_domain,
  jsonb_array_elements_text(nameservers) as ns;
```

```sql+sqlite
select
  domain,
  ns.value
from
  namecheap_domain,
  json_each(nameservers) as ns;
```