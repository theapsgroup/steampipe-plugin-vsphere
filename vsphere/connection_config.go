package vsphere

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type VsphereConfig struct {
	BaseUrl  *string `cty:"baseurl"`
	Insecure *bool   `cty:"insecure"`
	Username *string `cty:"username"`
	Password *string `cty:"password"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"baseurl": {
		Type: schema.TypeString,
	},
	"insecure": {
		Type: schema.TypeBool,
	},
	"username": {
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
