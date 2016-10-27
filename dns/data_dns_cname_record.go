package dns

import (
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataDnsCnameRecord() *schema.Resource {
	return &schema.Resource{
		Read: resourceDnsCnameRecordRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDnsCnameRecordRead(d *schema.ResourceData, meta interface{}) error {
	cname, err := net.LookupCNAME(d.Get("name").(string))
	if err != nil {
		return err
	}

	d.Set("cname", cname)
	return nil
}
