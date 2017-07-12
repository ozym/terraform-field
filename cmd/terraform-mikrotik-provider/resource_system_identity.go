package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikSystemIdentity() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikSystemIdentityCreate,
		Read:   resourceMikroTikSystemIdentityRead,
		Update: resourceMikroTikSystemIdentityUpdate,
		Delete: resourceMikroTikSystemIdentityDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMikroTikSystemIdentityCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikSystemIdentityUpdate(d, meta)
}

func resourceMikroTikSystemIdentityRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemIdentity()
		if err != nil {
			return err
		}

		if _, ok := res["name"]; ok {
			d.Set("name", res["name"])
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemIdentityUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if d.HasChange("name") {
			if err := mt.SetSystemIdentityName(d.Get("name").(string)); err != nil {
				return err
			}
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemIdentityDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		if err := mt.SetSystemIdentityName(""); err != nil {
			return err
		}

		return nil
	})
}
