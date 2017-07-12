package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingBgpAggregate() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikRoutingBgpAggregateCreate,
		Read:   resourceMikroTikRoutingBgpAggregateRead,
		Update: resourceMikroTikRoutingBgpAggregateUpdate,
		Delete: resourceMikroTikRoutingBgpAggregateDelete,

		Schema: map[string]*schema.Schema{
			"instance": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"include_igp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"summary_only": &schema.Schema{
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

func resourceMikroTikRoutingBgpAggregateCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikRoutingBgpAggregateUpdate(d, meta)
}

func resourceMikroTikRoutingBgpAggregateRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingBgpAggregate(d.Get("instance").(string), d.Get("prefix").(string))
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["include-igp"]; ok {
			d.Set("include_igp", ros.ParseBool(res["include-igp"]))
		}

		if _, ok := res["summary-only"]; ok {
			d.Set("summary_only", ros.ParseBool(res["summary-only"]))
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("instance").(string) + ":::" + d.Get("prefix").(string))

		return nil
	})
}

func resourceMikroTikRoutingBgpAggregateUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetRoutingBgpAggregateComment(d.Get("instance").(string), d.Get("prefix").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		if d.HasChange("include_igp") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_aggregate.include_igp parameter on %s", d.Get("name").(string))
		}
		if d.HasChange("summary_only") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_aggregate.summary_only parameter on %s", d.Get("name").(string))
		}
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_bgp_aggregate.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("instance").(string) + ":::" + d.Get("prefix").(string))

		return nil
	})
}

func resourceMikroTikRoutingBgpAggregateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
