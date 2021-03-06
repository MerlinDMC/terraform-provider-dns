package main

import (
	"github.com/hashicorp/terraform/plugin"

	"github.com/MerlinDMC/terraform-provider-dns/dns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: dns.Provider,
	})
}
