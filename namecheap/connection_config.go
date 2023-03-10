package namecheap

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type namecheapConfig struct {
	Username *string `cty:"username" hcl:"username"`
	APIUser  *string `cty:"api_user" hcl:"api_user"`
	APIKey   *string `cty:"api_key" hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &namecheapConfig{}
}

func GetConfig(connection *plugin.Connection) namecheapConfig {
	if connection == nil || connection.Config == nil {
		return namecheapConfig{}
	}
	config, _ := connection.Config.(namecheapConfig)
	return config
}
