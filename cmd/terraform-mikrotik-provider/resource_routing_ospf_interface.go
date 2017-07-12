package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingOspfInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikRoutingOspfInterfaceCreate,
		Read:   resourceMikroTikRoutingOspfInterfaceRead,
		Update: resourceMikroTikRoutingOspfInterfaceUpdate,
		Delete: resourceMikroTikRoutingOspfInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost": &schema.Schema{
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

func resourceMikroTikRoutingOspfInterfaceCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikRoutingOspfInterfaceUpdate(d, meta)
}

func resourceMikroTikRoutingOspfInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingOspfInterface(d.Get("interface").(string))
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["network-type"]; ok {
			d.Set("network_type", res["network-type"])
		}

		if _, ok := res["hello-interval"]; ok {
			d.Set("hello_interval", res["hello-interval"])
		}

		if _, ok := res["dead-interval"]; ok {
			d.Set("dead_interval", res["dead-interval"])
		}

		if _, ok := res["authentication"]; ok {
			d.Set("authentication", res["authentication"])
		}

		if _, ok := res["authentication-key"]; ok {
			d.Set("authentication_key", res["authentication-key"])
		}

		if _, ok := res["cost"]; ok {
			d.Set("cost", res["cost"])
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("interface").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfInterfaceUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetRoutingOspfInterfaceComment(d.Get("interface").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		/*
			if d.HasChange("cost") {
				if err := mt.SetRoutingOspfInterfaceComment(d.Get("interface").(string), d.Get("cost").(string)); err != nil {
					return err
				}
				d.SetPartial("cost")
			}
		*/

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("interface").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
