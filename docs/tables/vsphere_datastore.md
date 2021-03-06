# Table: vsphere_datastore

A datastore is a storage pool that can be used by virtual machines.

The `vsphere_datastore` table can be used to query datastore utilization and capacity.

## Examples

### List datastores

```sql
select
  *
from
  vsphere_datastore;
```

### Select inacessible datastores

```sql
select
  *
from
  vsphere_datastore
where
  accessible = false;
```

### Select NFS type datastores

```sql
select
  *
from
  vsphere_datastore
where
  type = 'NFS';
```