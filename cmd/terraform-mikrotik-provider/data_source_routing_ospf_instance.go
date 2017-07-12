package main

import (
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikRoutingOspfInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikRoutingOspfInstancesRead,

		Schema: map[string]*schema.Schema{
			"instances": &schema.Schema{
				Type:     schema.TypeSet,
				ForceNew: true,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"router_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"distribute_default": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redistribute_connected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redistribute_static": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redistribute_rip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redistribute_bgp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"redistribute_other_ospf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"metric_default": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"metric_connected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"metric_static": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"metric_rip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"metric_bgp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"metric_other_ospf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"in_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"out_filter": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
						"default": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
					},
				},
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					return hashcode.String(m["name"].(string))
				},
			},
		},
	}
}

func dataSourceMikroTikRoutingOspfInstancesRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingOspfInstances()
		if err != nil {
			return err
		}

		instances := []map[string]interface{}{}

		for _, r := range res {

			instance := map[string]interface{}{}

			for k, v := range map[string]string{
				"name":                    "name",
				"router_id":               "router-id",
				"distribute_default":      "distribute-default",
				"redistribute_connected":  "redistribute-connected",
				"redistribute_static":     "redistribute-static",
				"redistribute_rip":        "redistribute-rip",
				"redistribute_bgp":        "redistribute-bgp",
				"redistribute_other_ospf": "redistribute-other-ospf",
				"metric_default":          "metric-default",
				"metric_connected":        "metric-connected",
				"metric_static":           "metric-static",
				"metric_rip":              "metric-rip",
				"metric_bgp":              "metric-bgp",
				"metric_other_ospf":       "metric-other-ospf",
				"in_filter":               "in-filter",
				"out_filter":              "out-filter",
			} {
				if x, ok := r[v]; ok {
					instance[k] = x
				}
			}

			for k, v := range map[string]string{
				"disabled": "disabled",
				"default":  "default",
			} {
				if x, ok := r[v]; ok {
					instance[k] = ros.ParseBool(x)
				}
			}

			instances = append(instances, instance)
		}

		if err := d.Set("instances", instances); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}
