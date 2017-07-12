package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/ozym/airos"
)

// Provider returns a schema.Provider for Example.
func ubiquitiProvider() terraform.ResourceProvider {
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
			"ubiquiti_radio_status":   dataSourceUbiquitiRadioStatus(),
			"ubiquiti_radio_wireless": dataSourceUbiquitiRadioWireless(),
			"ubiquiti_radio_stations": dataSourceUbiquitiRadioStations(),
		},

		ConfigureFunc: ubiquitiConfigure,
	}
}

func ubiquitiConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname := d.Get("hostname").(string)
	port := d.Get("port").(int)

	community := d.Get("community").(string)

	retries := d.Get("retries").(int)
	timeout, err := time.ParseDuration(d.Get("timeout").(string))
	if err != nil {
		return nil, err
	}

	options := []func(*airos.AirOs) error{
		airos.Port(port), airos.Community(community), airos.Retries(retries), airos.Timeout(timeout),
	}

	return airos.NewAirOs(hostname, options...)
}
