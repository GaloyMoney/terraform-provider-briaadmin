package main

import (
	"github.com/GaloyMoney/terraform-provider-bria/pkg/admin"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: admin.Provider,
	})
}