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

## Getting Started

### Installation

Download and install the latest vSphere plugin:

```shell
steampipe plugin install theapsgroup/vsphere
```

### Prerequisites

- vSphere cluster
- (Readonly) credentials

### Configuration

Configuration can be done using both a configuration file and environment variables.

Note: Configuration file variables will take precedence over environment variables.

Configuration File:

```hcl
connection "vsphere" {
  plugin = "theapsgroup/vsphere"

  vsphere_server  = "192.168.122.233"
  user  = "root"
  password  = "s0Mep@ss"
  allow_unverified_ssl = true
}
```

Environment variables:

```
export VSPHERE_SERVER=192.168.122.233
export VSPHERE_USER=root
export VSPHERE_PASSWORD='s0Mep@ss'
export ALLOW_UNVERIFIED_SSL=true
```