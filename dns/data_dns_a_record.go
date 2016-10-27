package dns

import (
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataDnsARecord() *schema.Resource {
	return &schema.Resource{
		Read: resourceDnsARecordRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"addrs": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func resourceDnsARecordRead(d *schema.ResourceData, meta interface{}) error {
	addrs, err := net.LookupHost(d.Get("name").(string))
	if err != nil {
		return err
	}

	d.Set("addrs", addrs)
	return nil
}
