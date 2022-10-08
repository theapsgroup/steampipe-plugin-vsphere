package vsphere

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type VsphereConfig struct {
	VsphereServer      *string `cty:"vsphere_server"`
	AllowUnverifiedSSL *bool   `cty:"allow_unverified_ssl"`
	User               *string `cty:"user"`
	Password           *string `cty:"password"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"vsphere_server": {
		Type: schema.TypeString,
	},
	"allow_unverified_ssl": {
		Type: schema.TypeBool,
	},
	"user": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &VsphereConfig{}
}

func GetConfig(connection *plugin.Connection) VsphereConfig {
	if connection == nil || connection.Config == nil {
		return VsphereConfig{}
	}

	config, _ := connection.Config.(VsphereConfig)
	return config
}
