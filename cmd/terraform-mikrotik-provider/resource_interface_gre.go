package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikInterfaceGre() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikInterfaceGreCreate,
		Read:   resourceMikroTikInterfaceGreRead,
		Update: resourceMikroTikInterfaceGreUpdate,
		Delete: resourceMikroTikInterfaceGreDelete,

		Schema: map[string]*schema.Schema{
			"remote_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"local_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "auto",
			},
			"keepalive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s,10",
			},
			"clamp_tcp_mss": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"dont_fragment": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"allow_fast_path": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"minor": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"major": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

func resourceMikroTikInterfaceGreHasFastPath(d *schema.ResourceData, mt *ros.Ros) bool {
	switch {
	case mt.Major() < 6:
		return false
	case mt.Major() > 6:
		return true
	case mt.Minor() < 33:
		return false
	default:
		return true
	}
}

func resourceMikroTikInterfaceGreCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikInterfaceGreUpdate(d, meta)
}

func resourceMikroTikInterfaceGreRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.InterfaceGre(d.Get("remote_address").(string))
		if err != nil {
			return err
		}

		if _, ok := res["local-address"]; ok {
			d.Set("local_address", res["local-address"])
		}

		if _, ok := res["name"]; ok {
			d.Set("name", res["name"])
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["mtu"]; ok {
			d.Set("mtu", res["mtu"])
		}
		if _, ok := res["keepalive"]; ok {
			d.Set("keepalive", res["keepalive"])
		}

		if _, ok := res["clamp-tcp-mss"]; ok {
			d.Set("clamp_tcp_mss", ros.ParseBool(res["clamp-tcp-mss"]))
		}

		if _, ok := res["dont-fragment"]; ok {
			d.Set("dont_fragment", ros.ParseBool(res["dont-fragment"]))
		}

		if _, ok := res["allow-fast-path"]; ok {
			d.Set("allow_fast_path", ros.ParseBool(res["allow-fast-path"]))
		}
		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("remote_address").(string))

		return nil
	})
}

func resourceMikroTikInterfaceGreUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("name") {
			if err := mt.SetInterfaceGreComment(d.Get("remote_address").(string), d.Get("name").(string)); err != nil {
				return err
			}
			d.SetPartial("name")
		}
		if d.HasChange("comment") {
			if err := mt.SetInterfaceGreComment(d.Get("remote_address").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}
		if d.HasChange("mtu") {
			if err := mt.SetInterfaceGreMtu(d.Get("remote_address").(string), d.Get("mtu").(string)); err != nil {
				return err
			}
			d.SetPartial("mtu")
		}
		if d.HasChange("keepalive") {
			if err := mt.SetInterfaceGreKeepalive(d.Get("remote_address").(string), d.Get("keepalive").(string)); err != nil {
				return err
			}
			d.SetPartial("keepalive")
		}
		if d.HasChange("clamp_tcp_mss") {
			if err := mt.SetInterfaceGreClampTcpMss(d.Get("remote_address").(string), d.Get("clamp_tcp_mss").(bool)); err != nil {
				return err
			}
			d.SetPartial("clamp_tcp_mss")
		}
		if d.HasChange("dont_fragment") {
			if err := mt.SetInterfaceGreClampTcpMss(d.Get("remote_address").(string), d.Get("dont_fragment").(bool)); err != nil {
				return err
			}
			d.SetPartial("dont_fragment")
		}
		if d.HasChange("allow_fast_path") {
			if resourceMikroTikInterfaceGreHasFastPath(d, mt) {
				if err := mt.SetInterfaceGreAllowFastPath(d.Get("remote_address").(string), d.Get("allow_fast_path").(bool)); err != nil {
					return err
				}
			}
			d.SetPartial("allow_fast_path")
		}

		// read-only parameters
		if d.HasChange("local_address") {
			return fmt.Errorf("unable to set read-only mikrotik_interface_gre.local_address parameter on %s", d.Get("name").(string))
			//d.SetPartial("local_address")
		}
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_interface_gre.disabled parameter on %s", d.Get("name").(string))
			//d.SetPartial("disabled")
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("remote_address").(string))

		return nil
	})
}

func resourceMikroTikInterfaceGreDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
