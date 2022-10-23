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

### Select all vms with a name containing test and an uptime of more than 1 hour

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

### Total Actual storage consumption per VM in Gigabytes and on how many/which datastores

```sql
with  
    alldisks as (
        select 
            jsonb_array_elements(storageconsumed) as disks, 
            name, 
            moref 
        from vc.vsphere_vm) 
    select 
        moref, 
        (array_agg(name))[1] as Name, 
        sum(disks['Committed']::bigint/(1024*1024*1024)) as UsageGB, 
        count(disks['Committed']) as DatastoresCount,
        string_agg(disks['Datastore']['Value']::text, ', ') as Datastores
    from alldisks 
    group by moref
```

### Show all virtual disks in order of size that are used by all the VMs with size and type information

```sql
with disks as (
    with devices as (
        select 
            name, 
            moref, 
            jsonb_array_elements(devices) as device     
        from 
            vc.vsphere_vm
        ) 
    select 
        moref,
        name,
        trim(both '"' from device['DeviceInfo']['Label']::text) as label,
        device['CapacityInKB']::bigint/(1024*1024) as sizeinGB,
        trim(both '"' from device['Backing']['FileName']::text) as filename,
        device['Backing']['ThinProvisioned']::boolean as thinprovisioned
    from 
        devices
)
    select 
       *
    from
        disks
    where
        label like '%Hard disk%'
    order by sizeinGB
```