# Table: namecheap_domain_host

List all domain host records managed in the Namecheap account.

## Examples

### List domain host records

```sql
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

```sql
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

```sql
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