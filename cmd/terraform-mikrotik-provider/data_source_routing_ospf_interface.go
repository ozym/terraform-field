package main

import (
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikRoutingOspfInterfaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikRoutingOspfInterfacesRead,

		Schema: map[string]*schema.Schema{
			"interfaces": &schema.Schema{
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
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"cost": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"authentication_key": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"authentication_key_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"network_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"retransmit_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"transmit_delay": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dead_interval": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
						"dynamic": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
						"passive": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"inactive": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
						"use_bdf": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
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

func dataSourceMikroTikRoutingOspfInterfacesRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingOspfInterfaces()
		if err != nil {
			return err
		}

		interfaces := []map[string]interface{}{}

		for _, r := range res {

			iface := map[string]interface{}{}

			for k, v := range map[string]string{
				"interface":             "interface",
				"cost":                  "cost",
				"priority":              "priority",
				"authentication":        "authentication",
				"authentication_key":    "authentication-key",
				"authentication_key_id": "authentication-key-id",
				"network_type":          "network-type",
				"retransmit_interval":   "retransmit-interval",
				"transmit_delay":        "transmit-delay",
				"hello_interval":        "hello-interval",
				"dead_interval":         "dead-interval",
			} {
				if x, ok := r[v]; ok {
					iface[k] = x
				}
			}

			for k, v := range map[string]string{
				"disabled": "disabled",
				"dynamic":  "dynamic",
				"passive":  "passive",
				"inactive": "inactive",
				"use_bdf":  "use-bdf",
			} {
				if x, ok := r[v]; ok {
					iface[k] = ros.ParseBool(x)
				}
			}

			interfaces = append(interfaces, iface)
		}

		if err := d.Set("interfaces", interfaces); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}
