package dns

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		DataSourcesMap: map[string]*schema.Resource{
			"dns_a_record":     dataDnsARecord(),
			"dns_cname_record": dataDnsCnameRecord(),
			"dns_srv_record":   dataDnsSrvRecord(),
			"dns_txt_record":   dataDnsTxtRecord(),
		},
	}
}
