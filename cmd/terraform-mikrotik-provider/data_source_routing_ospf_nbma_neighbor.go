package main

import (
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikRoutingOspfNbmaNeighbors() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikRoutingOspfNbmaNeighborsRead,

		Schema: map[string]*schema.Schema{
			"neighbors": &schema.Schema{
				Type:     schema.TypeSet,
				ForceNew: true,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"instance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"poll_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
					},
				},
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					return hashcode.String(m["address"].(string))
				},
			},
		},
	}
}

func dataSourceMikroTikRoutingOspfNbmaNeighborsRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingOspfNbmaNeighbors()
		if err != nil {
			return err
		}

		neighbors := []map[string]interface{}{}

		for _, r := range res {

			neighbor := map[string]interface{}{}

			for k, v := range map[string]string{
				"address":       "address",
				"instance":      "instance",
				"poll_interval": "poll-interval",
				"priority":      "priority",
				"comment":       "comment",
			} {
				if x, ok := r[v]; ok {
					neighbor[k] = x
				}
			}

			for k, v := range map[string]string{
				"disabled": "disabled",
			} {
				if x, ok := r[v]; ok {
					neighbor[k] = ros.ParseBool(x)
				}
			}

			neighbors = append(neighbors, neighbor)
		}

		if err := d.Set("neighbors", neighbors); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}
