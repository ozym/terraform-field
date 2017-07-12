package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/ozym/skyedge"
)

// Provider returns a schema.Provider for Example.
func skyEdgeProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
			},
			/*
				"retries": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  3,
				},
			*/
		},

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{
			"skyedge_status": dataSourceSkyEdgeStatus(),
		},

		ConfigureFunc: skyEdgeConfigure,
	}
}

func skyEdgeConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname := d.Get("hostname").(string)
	port := d.Get("port").(int)

	username := d.Get("username").(string)
	password := d.Get("password").(string)

	timeout, err := time.ParseDuration(d.Get("timeout").(string))
	if err != nil {
		return nil, err
	}
	options := []func(*skyedge.SkyEdge) error{
		skyedge.Port(port), skyedge.Username(username), skyedge.Password(password), skyedge.Timeout(timeout),
	}

	return skyedge.NewSkyEdge(hostname, options...)
}
