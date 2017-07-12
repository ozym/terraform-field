package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikIpRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikIpRouteCreate,
		Read:   resourceMikroTikIpRouteRead,
		Update: resourceMikroTikIpRouteUpdate,
		Delete: resourceMikroTikIpRouteDelete,

		Schema: map[string]*schema.Schema{
			"dst_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway": &schema.Schema{
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

func resourceMikroTikIpRouteCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikIpRouteUpdate(d, meta)
}

func resourceMikroTikIpRouteRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.IpRoute(d.Get("dst_address").(string))
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", ros.ParseBool(res["comment"]))
		}

		if _, ok := res["gateway"]; ok {
			d.Set("gateway", res["gateway"])
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("dst_address").(string))

		return nil
	})
}

func resourceMikroTikIpRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetIpRouteComment(d.Get("dst_address").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		// read-only
		if d.HasChange("gateway") {
			return fmt.Errorf("unable to set read-only mikrotik_ip_address.gateway parameter on %s", d.Get("name").(string))
		}
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_ip_address.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("dst_address").(string))

		return nil
	})
}

func resourceMikroTikIpRouteDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
