package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingBgpInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikRoutingBgpInstanceCreate,
		Read:   resourceMikroTikRoutingBgpInstanceRead,
		Update: resourceMikroTikRoutingBgpInstanceUpdate,
		Delete: resourceMikroTikRoutingBgpInstanceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"as": &schema.Schema{
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
			"client_to_client_reflection": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceMikroTikRoutingBgpInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikRoutingBgpInstanceRead(d, meta)
}

func resourceMikroTikRoutingBgpInstanceRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingBgpInstance(d.Get("name").(string))
		if err != nil {
			return err
		}

		if _, ok := res["as"]; ok {
			d.Set("as", res["as"])
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["router-id"]; ok {
			d.Set("router_id", res["router-id"])
		}

		if _, ok := res["client_to_client_reflection"]; ok {
			d.Set("client_to_client_reflection", ros.ParseBool(res["client-to-client-reflection"]))
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikRoutingBgpInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("router_id") {
			if err := mt.SetRoutingBgpInstanceRouterId(d.Get("name").(string), d.Get("router_id").(string)); err != nil {
				return err
			}
			d.SetPartial("router_id")
		}

		if d.HasChange("comment") {
			if err := mt.SetRoutingBgpInstanceComment(d.Get("name").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		if d.HasChange("as") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_instance.as parameter on %s", d.Get("name").(string))
		}
		if d.HasChange("client_to_client_reflection") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_instance.client_to_client_reflection parameter on %s", d.Get("name").(string))
		}
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_instance.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikRoutingBgpInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
