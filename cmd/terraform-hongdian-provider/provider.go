package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/ozym/hongdian"
)

// Provider returns a schema.Provider for Example.
func hongdianProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  23,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "30s",
			},
		},

		ResourcesMap: map[string]*schema.Resource{},

		DataSourcesMap: map[string]*schema.Resource{
			"hongdian_status": dataSourceHongdianStatus(),
		},

		ConfigureFunc: hongdianConfigure,
	}
}

func hongdianConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname := d.Get("hostname").(string)
	port := d.Get("port").(int)

	password := d.Get("password").(string)

	timeout, err := time.ParseDuration(d.Get("timeout").(string))
	if err != nil {
		return nil, err
	}
	options := []func(*hongdian.Hongdian) error{
		hongdian.Port(port), hongdian.Password(password), hongdian.Timeout(timeout),
	}

	return hongdian.NewHongdian(hostname, options...)
}
