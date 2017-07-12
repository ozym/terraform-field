package main

import (
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikRoutingOspfNetworks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikRoutingOspfNetworksRead,

		Schema: map[string]*schema.Schema{
			"networks": &schema.Schema{
				Type:     schema.TypeSet,
				ForceNew: true,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"area": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
						"invalid": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
					},
				},
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					return hashcode.String(m["network"].(string))
				},
			},
		},
	}
}

func dataSourceMikroTikRoutingOspfNetworksRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.RoutingOspfNetworks()
		if err != nil {
			return err
		}

		networks := []map[string]interface{}{}

		for _, r := range res {

			network := map[string]interface{}{}

			for k, v := range map[string]string{
				"network": "network",
				"area":    "area",
				"comment": "comment",
			} {
				if x, ok := r[v]; ok {
					network[k] = x
				}
			}

			for k, v := range map[string]string{
				"disabled": "disabled",
				"invalid":  "invalid",
			} {
				if x, ok := r[v]; ok {
					network[k] = ros.ParseBool(x)
				}
			}

			networks = append(networks, network)
		}

		if err := d.Set("networks", networks); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}
