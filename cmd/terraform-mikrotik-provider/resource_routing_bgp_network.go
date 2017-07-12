package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingBgpNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikRoutingBgpNetworkCreate,
		Read:   resourceMikroTikRoutingBgpNetworkRead,
		Update: resourceMikroTikRoutingBgpNetworkUpdate,
		Delete: resourceMikroTikRoutingBgpNetworkDelete,

		Schema: map[string]*schema.Schema{
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"synchronize": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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

func resourceMikroTikRoutingBgpNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikRoutingBgpNetworkUpdate(d, meta)
}

func resourceMikroTikRoutingBgpNetworkRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingBgpNetwork(d.Get("network").(string))
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["synchronize"]; ok {
			d.Set("synchronize", ros.ParseBool(res["synchronize"]))
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("network").(string))

		return nil
	})
}

func resourceMikroTikRoutingBgpNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetRoutingBgpNetworkComment(d.Get("network").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		if d.HasChange("synchronize") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_network.synchronize parameter on %s", d.Get("name").(string))
		}

		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_network.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("network").(string))

		return nil
	})
}

func resourceMikroTikRoutingBgpNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
