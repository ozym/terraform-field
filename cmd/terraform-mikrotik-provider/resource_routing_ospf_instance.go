package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingOspfInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikRoutingOspfInstanceCreate,
		Read:   resourceMikroTikRoutingOspfInstanceRead,
		Update: resourceMikroTikRoutingOspfInstanceUpdate,
		Delete: resourceMikroTikRoutingOspfInstanceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0/0",
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"redistribute_connected": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"redistribute_static": &schema.Schema{
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

func resourceMikroTikRoutingOspfInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikRoutingOspfInstanceUpdate(d, meta)
}

func resourceMikroTikRoutingOspfInstanceRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingOspfInstance(d.Get("name").(string))
		if err != nil {
			return err
		}

		if _, ok := res["router-id"]; ok {
			d.Set("router_id", res["router-id"])
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["redistribute-connected"]; ok {
			d.Set("redistribute_connected", res["redistribute-connected"])
		}

		if _, ok := res["redistribute-static"]; ok {
			d.Set("redistribute_static", res["redistribute-static"])
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("router_id") {
			if err := mt.SetRoutingOspfInstanceRouterId(d.Get("name").(string), d.Get("router_id").(string)); err != nil {
				return err
			}
			d.SetPartial("router_id")
		}

		if d.HasChange("comment") {
			if err := mt.SetRoutingOspfInstanceComment(d.Get("name").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
