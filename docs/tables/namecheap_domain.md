# Table: namecheap_domain

List all domains registered in the Namecheap account.

## Examples

### List domains

```sql
select
  domain,
  expires
from
  namecheap_domain;
```

### Get a specific domain

```sql
select
  domain,
  expires
from
  namecheap_domain
where
  domain = 'steampipe.io';
```

### Domains expiring in the next 90 days

```sql
select
  domain,
  expires,
  age(expires)
from
  namecheap_domain
where
  expires < (current_date + interval '90 days');
```

### Domains by date created at Namecheap

```sql
select
  domain,
  created,
  age(created)
from
  namecheap_domain
order by
  created;
```

### Domains without auto-renew enabled

```sql
select
  domain,
  expires
from
  namecheap_domain
where
  not auto_renew;
```

### Domains and their nameservers

```sql
select
  domain,
  ns
from
  namecheap_domain,
  jsonb_array_elements_text(nameservers) as ns;
```
