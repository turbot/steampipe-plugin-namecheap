package main

import (
	"github.com/turbot/steampipe-plugin-namecheap/namecheap"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: namecheap.Plugin})
}
