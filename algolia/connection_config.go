package algolia

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type algoliaConfig struct {
	AppID  *string `cty:"app_id"`
	APIKey *string `cty:"api_key"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"app_id": {
		Type: schema.TypeString,
	},
	"api_key": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &algoliaConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) algoliaConfig {
	if connection == nil || connection.Config == nil {
		return algoliaConfig{}
	}
	config, _ := connection.Config.(algoliaConfig)
	return config
}
