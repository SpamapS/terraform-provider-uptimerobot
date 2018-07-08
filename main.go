package main

import (
		"github.com/hashicorp/terraform/plugin"
		"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServOpts{
			ProviderFunc: func() terraform.ResourceProvider {
					return Provider()
			},
	})
}
