package main

import (
	"fmt"
	"strconv"
	//	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikRoutingFilterChain() *schema.Resource {
	return &schema.Resource{

		Create: resourceMikroTikRoutingFilterChainCreate,
		Read:   resourceMikroTikRoutingFilterChainRead,
		Update: resourceMikroTikRoutingFilterChainUpdate,
		Delete: resourceMikroTikRoutingFilterChainDelete,

		Schema: map[string]*schema.Schema{
			"chain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"prefix_length": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"set_route_comment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"set_bgp_local_pref": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"set_bgp_med": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"match_chain": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"invert_match": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"disabled": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		},
	}

}

func resourceMikroTikRoutingFilterChainCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikRoutingFilterChainUpdate(d, meta)
}

func resourceMikroTikRoutingFilterChainRead(d *schema.ResourceData, meta interface{}) error {

	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		list, err := mt.RoutingFilterChain(d.Get("chain").(string))
		if err != nil {
			return err
		}

		/*
			if n, ok := d.GetOk("rules.#"); ok {
				if len(list) != n {
					return fmt.Errorf("%s: incorrect number of mikrotik_routing_filter_chain.rules.# found, expected %d, read %d", d.Get("chain").(string), n, len(list))
				}
			}
		*/

		d.Set("rules.#", len(list))

		for num, res := range list {
			if action, ok := res["action"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".action", action)
			}
			if prefix, ok := res["prefix"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".prefix", prefix)
			}
			if length, ok := res["prefix-length"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".prefix_length", length)
			}
			if protocol, ok := res["protocol"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".protocol", protocol)
			}
			if comment, ok := res["set-route-comment"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".set_route_comment", comment)
			}
			if pref, ok := res["set-bgp-local-pref"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".set_bgp_local_pref", pref)
			}
			if med, ok := res["set-bgp-med"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".set_bgp_med", med)
			}
			if comment, ok := res["comment"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".comment", comment)
			}
			if chain, ok := res["match-chain"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".match_chain", chain)
			}
			if match, ok := res["invert-match"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".invert_match", match)
			}
			if disabled, ok := res["disabled"]; ok {
				d.Set("rules."+strconv.Itoa(num)+".disabled", disabled)
			}
		}

		d.SetId(mt.Id() + ":::" + d.Get("chain").(string))

		return nil
	})
}

func resourceMikroTikRoutingFilterChainUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if d.HasChange("rules.#") {
			return fmt.Errorf("unable to set read-only mikrotik_routing_filter_chain.rules length on %s", d.Get("chain").(string))
		}

		if rules, ok := d.GetOk("rules.#"); ok {
			for num := 0; num < rules.(int); num++ {
				if d.HasChange("rules." + strconv.Itoa(num) + ".comment") {
					//
				}

				// read-only
				if d.HasChange("rules." + strconv.Itoa(num) + ".disabled") {
					return fmt.Errorf("unable to set read-only mikrotik_routing_filter_chain.rules.%d.disabled parameter on %s", num, d.Get("chain").(string))
				}
			}
		}

		d.SetId(mt.Id() + ":::" + d.Get("chain").(string))

		return nil
	})
}

func resourceMikroTikRoutingFilterChainDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
