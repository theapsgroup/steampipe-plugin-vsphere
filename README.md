# vSphere plugin for Steampipe

[Steampipe](https://steampipe.io) plugin for vsphere.

## Query vSphere with Steampipe

```sql
select * from vsphere_vm;
```

## Get Started

### Installation

```shell
steampipe plugin install theapsgroup/vsphere
```

Or if you prefer, you can clone this repository and build/install from source directly.

```shell
go build -o steampipe-plugin-vsphere.plugin

mv steampipe-plugin-vsphere.plugin ~/.steampipe/plugins/hub.steampipe.io/plugins/theapsgroup/vsphere@latest/steampipe-plugin-vsphere.plugin

cp config/vsphere.spc ~/.steampipe/config/vsphere.spc
```

Configuration is done via the configuration file:
`vi ~/.steampipe/config/vsphere.spc` 

## Documentation

Further documentation can he [found here](https://github.com/theapsgroup/steampipe-plugin-vsphere/blob/main/docs/index.md)

## Credits

A Go library for interacting with VMware vSphere APIs [govmomi](https://github.com/vmware/govmomi) (licensed separately using this [Apache License](https://github.com/vmware/govmomi/blob/master/LICENSE.txt))
