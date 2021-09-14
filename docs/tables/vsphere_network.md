# Table: vsphere_network

A network is an isolated virtual network that vms can use to communicate between each other or to the internet.

The `vsphere_network` table can be used to query networks and availability.

## Examples

### List networks

```sql
select
  *
from
  vsphere_network;
```