package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikVersion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikVersionRead,

		Schema: map[string]*schema.Schema{
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"major": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"minor": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceMikroTikVersionRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		version, versionOk := d.GetOk("version")
		if !versionOk {
			return fmt.Errorf("no version found")
		}

		major, minor := ros.RouterOsVersion(version.(string))

		d.Set("major", major)
		d.Set("minor", minor)

		d.SetId(mt.Id())

		return nil
	})
}
