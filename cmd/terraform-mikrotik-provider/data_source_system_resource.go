package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikSystemResource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikSystemResourceRead,

		Schema: map[string]*schema.Schema{
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"architecture_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"board_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceMikroTikSystemResourceRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemResource()
		if err != nil {
			return err
		}

		for k, v := range map[string]string{
			"version":           "version",
			"cpu":               "cpu",
			"architecture_name": "architecture-name",
			"board_name":        "board-name",
			"platform":          "platform",
		} {
			if x, ok := res[v]; ok {
				d.Set(k, x)
			}
		}

		d.SetId(mt.Id())

		return nil
	})
}
