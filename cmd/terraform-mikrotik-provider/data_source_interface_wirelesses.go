package main

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func dataSourceMikroTikInterfaceWirelesses() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMikroTikInterfaceWirelessesRead,

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
						"interface_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"ssid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"band": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"channel_width": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wireless_protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wds_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"wds_default_bridge": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"bridge_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"frequency": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"l2mtu": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"running": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
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

func dataSourceMikroTikInterfaceWirelessesRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if mt.Major() > 0 && mt.Major() < 6 {
			d.SetId(mt.Id())
			return nil
		}

		res, err := mt.InterfaceWirelesses()
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
				"name":               "name",
				"comment":            "comment",
				"interface_type":     "interface-type",
				"mode":               "mode",
				"ssid":               "ssid",
				"mac_address":        "mac-address",
				"band":               "band",
				"channel_width":      "channel-width",
				"wireless_protocol":  "wireless-protocol",
				"wds_mode":           "wds-mode",
				"wds_default_bridge": "wds-default-bridge",
				"bridge_mode":        "bridge-mode",
			} {
				if x, ok := r[v]; ok {
					in[k] = x
				}
			}

			for k, v := range map[string]string{
				"mtu":       "mtu",
				"l2mtu":     "l2mtu",
				"frequency": "frequency",
			} {
				if x, ok := r[v]; ok {
					if n, err := strconv.Atoi(x); err == nil {
						in[k] = n
					}
				}
			}

			for k, v := range map[string]string{
				"disabled": "disabled",
				"running":  "running",
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

		return nil
	})
}
