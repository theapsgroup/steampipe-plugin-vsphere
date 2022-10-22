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

### Select all VMs showing their powerstate and on which host they are running

```sql
 select 
    vm.name, host.name, vm.power 
  from vc.vsphere_vm as vm 
  inner join vc.vsphere_host as host 
  on vm.hostmoref = host.hostmoref
```

### Total Actual disk usage per VM 

```sql
with  
    alldisks as (
        select 
            jsonb_array_elements(storage) as disks, 
            name, 
            moref 
        from vc.vsphere_vm) 
    select 
        moref, 
        (array_agg(name))[1] as Name, 
        sum(disks['Committed']::bigint/(1024*1024*1024)) as UsageGB, 
        count(disks['Committed']) as NumDisks 
    from alldisks 
    group by moref
```

