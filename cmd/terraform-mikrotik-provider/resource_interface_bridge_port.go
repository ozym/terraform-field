package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikInterfaceBridgePort() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikInterfaceBridgePortCreate,
		Read:   resourceMikroTikInterfaceBridgePortRead,
		Update: resourceMikroTikInterfaceBridgePortUpdate,
		Delete: resourceMikroTikInterfaceBridgePortDelete,

		Schema: map[string]*schema.Schema{
			"bridge": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceMikroTikInterfaceBridgePortCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikInterfaceBridgePortUpdate(d, meta)
}

func resourceMikroTikInterfaceBridgePortRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.InterfaceBridgePort(d.Get("bridge").(string), d.Get("interface").(string))
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("bridge").(string) + ":::" + d.Get("interface").(string))

		return nil
	})
}

func resourceMikroTikInterfaceBridgePortUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetInterfaceBridgePortComment(d.Get("bridge").(string), d.Get("interface").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		// read-only settings
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_interface_bridge_port.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("bridge").(string) + ":::" + d.Get("interface").(string))

		return nil
	})
}

func resourceMikroTikInterfaceBridgePortDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
