package main

import (
	"github.com/bgshacklett/terraform-provider-skeleton/skeleton"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: skeleton.Provider})
}
