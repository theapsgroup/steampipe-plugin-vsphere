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

Configuration File:

```hcl
connection "vsphere" {
  plugin = "theapsgroup/vsphere"

  baseurl  = "192.168.122.233"
  username  = "root"
  password  = "s0Mep@ss"
  insecure = true
}
```