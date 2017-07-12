package main

import (
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikIpAddresses() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikIpAddressesRead,

		Schema: map[string]*schema.Schema{
			"major": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"addresses": &schema.Schema{
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
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"actual_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dynamic": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
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
					return hashcode.String(m["address"].(string))
				},
			},
		},
	}
}

func dataSourceMikroTikIpAddressesRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if mt.Major() > 0 && mt.Major() < 6 {
			d.SetId(mt.Id())

			return nil
		}

		res, err := mt.IpAddresses()
		if err != nil {
			return err
		}

		addresses := []map[string]interface{}{}

		for _, r := range res {

			addr := map[string]interface{}{}

			for k, v := range map[string]string{
				"address":          "address",
				"comment":          "comment",
				"network":          "network",
				"interface":        "interface",
				"actual_interface": "actual-interface",
			} {
				if x, ok := r[v]; ok {
					addr[k] = x
				}
			}

			for k, v := range map[string]string{
				"dynamic":  "dynamic",
				"disabled": "disabled",
				"invalid":  "invalid",
			} {
				if x, ok := r[v]; ok {
					addr[k] = ros.ParseBool(x)
				}
			}

			addresses = append(addresses, addr)
		}

		if err := d.Set("addresses", addresses); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}
