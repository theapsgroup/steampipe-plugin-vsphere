# Table: vsphere_host

A host is a machine that provides the compute for virtual machines and other vSphere features.

The `vsphere_host` table can be used to query host utilization and hardware information.

## Examples

### List hosts

```sql
select
  *
from
  vsphere_host;
```