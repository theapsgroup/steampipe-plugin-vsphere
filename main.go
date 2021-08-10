package main

import (
	"steampipe-plugin-vsphere/vsphere"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: vsphere.Plugin})
}
