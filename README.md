![image](https://hub.steampipe.io/images/plugins/theapsgroup/vsphere-social-graphic.png)
# vSphere plugin for Steampipe

Use SQL to query information about your vSphere resources.

- **[Get started →](https://hub.steampipe.io/plugins/theapsgroup/vsphere)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/theapsgroup/vsphere/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/theapsgroup/steampipe-plugin-vsphere/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io/downloads):

```shell
steampipe plugin install theapsgroup/vsphere
```

[Configure the plugin](https://hub.steampipe.io/plugins/theapsgroup/vsphere#configuration) using the configuration file:

```shell
vi ~/.steampipe/vsphere.spc
```

Or environment variables:

```shell
export VSPHERE_SERVER=10.20.30.40
export VSPHERE_USER=bob
export VSPHERE_PASSWORD=s0m3p@ss
```

Start Steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
  name,
  num_cpu,
  ip_address
from
  vsphere_vm;
```

## Developing

Prerequisites:

* [Steampipe](https://steampipe.io/downloads)
* [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/theapsgroup/steampipe-plugin-vsphere.git
cd steampipe-plugin-vsphere
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/vsphere.spc
```

Try it!

```shell
steampipe query
> .inspect vsphere
```

Further reading:

* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-github/blob/main/LICENSE).

## Credits

A Go library for interacting with VMware vSphere APIs [govmomi](https://github.com/vmware/govmomi) (licensed separately using this [Apache License](https://github.com/vmware/govmomi/blob/master/LICENSE.txt))
