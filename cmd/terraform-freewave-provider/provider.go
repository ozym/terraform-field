package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/ozym/fgr"
)

// Provider returns a schema.Provider for Example.
func freewaveProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  161,
			},
			"community": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "public",
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
			},
			"retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3,
			},
		},

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{
			"freewave_radio_status": dataSourceFreewaveRadioStatus(),
		},

		ConfigureFunc: freewaveConfigure,
	}
}

func freewaveConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname := d.Get("hostname").(string)
	port := d.Get("port").(int)

	community := d.Get("community").(string)

	retries := d.Get("retries").(int)
	timeout, err := time.ParseDuration(d.Get("timeout").(string))
	if err != nil {
		return nil, err
	}

	options := []func(*fgr.Fgr) error{
		fgr.Port(port), fgr.Community(community), fgr.Retries(retries), fgr.Timeout(timeout),
	}

	return fgr.NewFgr(hostname, options...)
}
