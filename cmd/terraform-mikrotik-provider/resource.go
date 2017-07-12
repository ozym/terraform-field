package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTik() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikCreate,
		Read:   resourceMikroTikRead,
		Update: resourceMikroTikUpdate,
		Delete: resourceMikroTikDelete,

		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceMikroTikCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikUpdate(d, meta)
}

func resourceMikroTikRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		d.SetId(mt.Id())
		return nil
	})
}

func resourceMikroTikUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		d.SetId(mt.Id())
		return nil
	})
}

func resourceMikroTikDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		return nil
	})
}
