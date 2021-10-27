# Table: vsphere_vm

A VM is a virtual machine.

The `vsphere_vm` table can be used to query virtual machines.

## Examples

### List virtual machines

```sql
select
  *
from
  vsphere_vm;
```
### Select all vms with more than 6 cores assigned

```sql
select
  *
from
  vsphere_vm
where
  num_cpu > 6;
```

### Select all vms with a name containing test and and uptime of more than 1 hour

```sql
select
  *
from
  vsphere_vm
where
  name ILIKE '%test%' and uptime > 3600;
```