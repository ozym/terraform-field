package main

import (
	//	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

// manage non-structural elements of an existing ip address
func resourceMikroTikIpAddress() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikIpAddressCreate,
		Read:   resourceMikroTikIpAddressRead,
		Update: resourceMikroTikIpAddressUpdate,
		Delete: resourceMikroTikIpAddressDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceMikroTikIpAddressCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikIpAddressUpdate(d, meta)
}

func resourceMikroTikIpAddressRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.IpAddress(d.Get("address").(string))
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}
		if _, ok := res["interface"]; ok {
			d.Set("interface", res["interface"])
		}
		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("address").(string))

		return nil
	})
}

func resourceMikroTikIpAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetIpAddressComment(d.Get("address").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("address").(string))

		return nil
	})
}

func resourceMikroTikIpAddressDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
