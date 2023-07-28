---
organization: The APS Group
category: ["virtualization"]
icon_url: "images/plugins/theapsgroup/vsphere.svg"
brand_color: "#696566"
display_name: "VMware vSphere"
short_name: "vsphere"
description: "Steampipe plugin for querying data from a vsphere environment."
og_description: Query VMware vSphere with SQL! Open source CLI. No DB required.
og_image: "images/plugins/theapsgroup/vsphere-social-graphic.png"
---

# VMware vSphere + Turbot Steampipe

[vSphere](https://www.vmware.com/nl/products/vsphere.html) VMware vSphere is VMware's virtualization platform, which transforms data centers into aggregated computing infrastructures that include CPU, storage, and networking resources by [VMware](https://www.vmware.com/)

[Steampipe](https://steampipe.io/) is an open source CLI for querying cloud APIs using SQL from [Turbot](https://turbot.com/)

For example:
```sql
select
  name,
  num_cpu,
  ip_address
from
  vsphere_vm;
```
```
+---------+---------+-------------+
| name    | num_cpu | ip_address  |
+---------+---------+-------------+
| host007 | 12      | 10.20.42.42 |
+---------+---------+-------------+
```

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/theapsgroup/vsphere/tables)**

## Get Started

### Install

Download and install the latest vSphere plugin:

```shell
steampipe plugin install theapsgroup/vsphere
```

### Configuration

Installing the latest vSphere plugin will create a config file (`~/.steampipe/config/vsphere.spc`) with a single connection named `vsphere`:

```hcl
connection "vsphere" {
  plugin = "theapsgroup/vsphere"

  # The IP/url of vSphere server - can also be set with VSPHERE_SERVER environment variable
  # vsphere_server = "192.168.122.233"

  # vSphere username - can also be set with VSPHERE_USER environment variable
  # user = "administrator@vsphere.local"

  # vsphere password - can also be set with VSPHERE_PASSWORD environment variable
  # password  = "s0Mep@ss"

  # TLS cert validation - can also be set with VSPHERE_ALLOW_UNVERIFIED_SSL environment variable
  # allow_unverified_ssl = true
}
```

- `vsphere_server` - The ip address or url of your vSphere server, can also be set with VSPHERE_SERVER environment variable.
- `user` - The username used to authenticate with vSphere, can also be set with VSPHERE_USER environment variable.
- `password` - The password used to authenticate with vSphere, can also be set with VSPHERE_PASSWORD environment variable.
- `allow_unverified_ssl` - Indicates if TLS certificate validation is required or not, can also be set with VSPHERE_ALLOW_UNVERIFIED_SSL environment variable.


## Get involved

- Open source: https://github.com/theapsgroup/steampipe-plugin-vsphere
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
