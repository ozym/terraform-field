package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingBgpPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikRoutingBgpPeerCreate,
		Read:   resourceMikroTikRoutingBgpPeerRead,
		Update: resourceMikroTikRoutingBgpPeerUpdate,
		Delete: resourceMikroTikRoutingBgpPeerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_as": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func resourceMikroTikRoutingBgpPeerCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikRoutingBgpPeerUpdate(d, meta)
}

func resourceMikroTikRoutingBgpPeerRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingBgpPeer(d.Get("remote_address").(string))
		if err != nil {
			return err
		}

		if _, ok := res["name"]; ok {
			d.Set("name", res["name"])
		}

		if _, ok := res["remote-as"]; ok {
			d.Set("remote_as", res["remote-as"])
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikRoutingBgpPeerUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetRoutingBgpPeerComment(d.Get("remote_address").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		if d.HasChange("name") {
			if err := mt.SetRoutingBgpPeerName(d.Get("remote_address").(string), d.Get("name").(string)); err != nil {
				return err
			}
			d.SetPartial("name")
		}

		if d.HasChange("remote_as") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_peer.remote_as parameter on %s", d.Get("name").(string))
		}
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_peer.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)
		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikRoutingBgpPeerDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
