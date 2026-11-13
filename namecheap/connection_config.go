package namecheap

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type namecheapConfig struct {
	Username *string `hcl:"username"`
	APIUser  *string `hcl:"api_user"`
	APIKey   *string `hcl:"api_key"`
}

func ConfigInstance() interface{} {
	return &namecheapConfig{}
}

func GetConfig(connection *plugin.Connection) namecheapConfig {
	if connection == nil || connection.GetConfig() == nil {
		return namecheapConfig{}
	}
	config, _ := connection.GetConfig().(namecheapConfig)
	return config
}
