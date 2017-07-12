package main

import (
	//"fmt"
	"strconv"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikInterfaces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikInterfacesRead,

		Schema: map[string]*schema.Schema{
			"interfaces": &schema.Schema{
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
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"default_name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"actual_mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_l2mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"l2mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							//Optional: true,
						},
						"fast_path": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
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
						"active": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							//Optional: true,
						},
						"slave": &schema.Schema{
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

func dataSourceMikroTikInterfacesRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if mt.Major() > 0 && mt.Major() < 6 {
			d.SetId(mt.Id())

			return nil
		}

		res, err := mt.Interfaces()
		if err != nil {
			return err
		}

		interfaces := []map[string]interface{}{}

		for _, r := range res {
			if x, ok := r["active"]; ok {
				if active := ros.ParseBool(x); !active {
					continue
				}
			}

			in := map[string]interface{}{}

			for k, v := range map[string]string{
				"name":         "name",
				"comment":      "comment",
				"default_name": "default-name",
				"type":         "type",
				"mac_address":  "mac-address",
			} {
				if x, ok := r[v]; ok {
					in[k] = x
				}
			}

			for k, v := range map[string]string{
				"mtu":        "mtu",
				"l2mtu":      "l2mtu",
				"actual_mtu": "actual-mtu",
				"max_l2mtu":  "max-l2mtu",
			} {
				if x, ok := r[v]; ok {
					if n, err := strconv.Atoi(x); err == nil {
						in[k] = n
					}
				}
			}

			for k, v := range map[string]string{
				"fast_path": "fast-path",
				"dynamic":   "dynamic",
				"disabled":  "disabled",
				"active":    "active",
				"slave":     "slave",
			} {
				if x, ok := r[v]; ok {
					in[k] = ros.ParseBool(x)
				}
			}

			interfaces = append(interfaces, in)
		}

		if err := d.Set("interfaces", interfaces); err != nil {
			return err
		}

		d.SetId(mt.Id())

		//return fmt.Errorf("%v\n", interfaces)

		return nil
	})
}
