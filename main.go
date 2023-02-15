package main

import (
	"steampipe-plugin-vsphere/vsphere"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: vsphere.Plugin})
}
