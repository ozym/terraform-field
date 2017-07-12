package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingOspfNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikRoutingOspfNetworkCreate,
		Read:   resourceMikroTikRoutingOspfNetworkRead,
		Update: resourceMikroTikRoutingOspfNetworkUpdate,
		Delete: resourceMikroTikRoutingOspfNetworkDelete,

		Schema: map[string]*schema.Schema{
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"area": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "backbone",
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceMikroTikRoutingOspfNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikRoutingOspfNetworkUpdate(d, meta)
}

func resourceMikroTikRoutingOspfNetworkRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingOspfNetwork(d.Get("network").(string))
		if err != nil {
			return err
		}

		if _, ok := res["area"]; ok {
			d.Set("area", res["area"])
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("network").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetRoutingOspfNetworkComment(d.Get("network").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		// read-only
		if d.HasChange("area") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_ospf_network.area parameter on %s", d.Get("name").(string))
		}
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_ospf_network.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("network").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
