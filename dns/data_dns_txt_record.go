package dns

import (
	"net"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataDnsTxtRecord() *schema.Resource {
	return &schema.Resource{
		Read: resourceDnsTxtRecordRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"record": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"records": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func resourceDnsTxtRecordRead(d *schema.ResourceData, meta interface{}) error {
	records, err := net.LookupTXT(d.Get("name").(string))
	if err != nil {
		return err
	}

	if len(records) > 0 {
		d.Set("record", records[0])
	}
	d.Set("records", records)
	return nil
}
