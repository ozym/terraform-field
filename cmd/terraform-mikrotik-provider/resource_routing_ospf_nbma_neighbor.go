package main

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingOspfNbmaNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikRoutingOspfNbmaNeighborCreate,
		Read:   resourceMikroTikRoutingOspfNbmaNeighborRead,
		Update: resourceMikroTikRoutingOspfNbmaNeighborUpdate,
		Delete: resourceMikroTikRoutingOspfNbmaNeighborDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
			"poll_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "5s",
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceMikroTikRoutingOspfNbmaNeighborCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		if err := mt.SetRoutingOspfNbmaNeighborComment(d.Get("address").(string), d.Get("comment").(string)); err != nil {
			return err
		}

		d.SetId(mt.Id() + ":::" + d.Get("address").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfNbmaNeighborRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingOspfNbmaNeighbor(d.Get("address").(string))
		if err != nil {
			return err
		}

		if _, ok := res["address"]; ok {
			d.Set("address", res["address"])
		}

		if _, ok := res["instance"]; ok {
			d.Set("instance", res["instance"])
		}

		if _, ok := res["poll-interval"]; ok {
			d.Set("poll_interval", res["poll-interval"])
		}

		if _, ok := res["priority"]; ok {
			priority, err := strconv.Atoi(res["priority"])
			if err != nil {
				return err
			}
			d.Set("priority", priority)
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		d.SetId(mt.Id() + ":::" + d.Get("address").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfNbmaNeighborUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetRoutingOspfNbmaNeighborComment(d.Get("address").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		if d.HasChange("poll_interval") {
			if err := mt.SetRoutingOspfNbmaNeighborPollInterval(d.Get("address").(string), d.Get("poll_interval").(string)); err != nil {
				return err
			}
			d.SetPartial("poll_interval")
		}

		if d.HasChange("priority") {
			if err := mt.SetRoutingOspfNbmaNeighborPriority(d.Get("address").(string), d.Get("priority").(int)); err != nil {
				return err
			}
			d.SetPartial("priority")
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("address").(string))

		return nil
	})
}

func resourceMikroTikRoutingOspfNbmaNeighborDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
