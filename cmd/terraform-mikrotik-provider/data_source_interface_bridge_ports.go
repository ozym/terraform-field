package main

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikInterfaceBridgePorts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikInterfaceBridgePortsRead,

		Schema: map[string]*schema.Schema{
			"major": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"number": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},

			"ports": &schema.Schema{
				Type:     schema.TypeSet,
				ForceNew: true,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"bridge": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"path_cost": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"edge": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"point_to_point": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"external_fdb": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"horizon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"inactive": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"dynamic": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					return hashcode.String(m["interface"].(string))
				},
			},
		},
	}
}

func dataSourceMikroTikInterfaceBridgePortsRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if mt.Major() > 0 && mt.Major() < 6 {
			d.SetId(mt.Id())
			return nil
		}

		res, err := mt.InterfaceBridgePorts()
		if err != nil {
			return err
		}

		ports := []map[string]interface{}{}

		for _, r := range res {

			in := map[string]interface{}{}

			for k, v := range map[string]string{
				"interface":      "interface",
				"bridge":         "bridge",
				"priority":       "priority",
				"edge":           "edge",
				"point_to_point": "point-to-point",
				"external_fdb":   "external-fdb",
				"horizon":        "horizon",
			} {
				if x, ok := r[v]; ok {
					in[k] = x
				}
			}

			for k, v := range map[string]string{
				"path_cost": "path-cost",
			} {
				if x, ok := r[v]; ok {
					if n, err := strconv.Atoi(x); err == nil {
						in[k] = n
					}
				}
			}

			for k, v := range map[string]string{
				"disabled": "disabled",
				"inactive": "inactive",
				"dynamic":  "dynamic",
			} {
				if x, ok := r[v]; ok {
					in[k] = ros.ParseBool(x)
				}
			}

			ports = append(ports, in)
		}

		if err := d.Set("ports", ports); err != nil {
			return err
		}

		if err := d.Set("number", len(ports)); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}
