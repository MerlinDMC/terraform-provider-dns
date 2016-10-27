package dns

import (
	"net"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataDnsSrvRecord() *schema.Resource {
	return &schema.Resource{
		Read: resourceDnsSrvRecordRead,

		Schema: map[string]*schema.Schema{
			"service": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				ForceNew: true,
			},
			"proto": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"hosts": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"addrs": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func resourceDnsSrvRecordRead(d *schema.ResourceData, meta interface{}) error {
	cname, addrs, err := net.LookupSRV(d.Get("service").(string), d.Get("proto").(string), d.Get("name").(string))
	if err != nil {
		return err
	}

	d.SetId(cname)

	var o_hosts []string
	var o_addrs []string

	for _, addr := range addrs {
		o_hosts = append(o_hosts, strings.TrimSuffix(addr.Target, "."))
		o_addrs = append(o_addrs, formatDnsSrvAddr(addr))
	}

	d.Set("hosts", o_hosts)
	d.Set("addrs", o_addrs)

	return nil
}

func formatDnsSrvAddr(addr *net.SRV) string {
	target := strings.TrimSuffix(addr.Target, ".")
	if addr.Port != 0 {
		return target + ":" + strconv.Itoa((int)(addr.Port))
	}
	return target
}
