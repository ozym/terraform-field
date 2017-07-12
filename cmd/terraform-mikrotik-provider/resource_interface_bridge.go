package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikInterfaceBridge() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikInterfaceBridgeCreate,
		Read:   resourceMikroTikInterfaceBridgeRead,
		Update: resourceMikroTikInterfaceBridgeUpdate,
		Delete: resourceMikroTikInterfaceBridgeDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "rstp",
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0x8000",
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

func resourceMikroTikInterfaceBridgeCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikInterfaceBridgeUpdate(d, meta)
}

func resourceMikroTikInterfaceBridgeRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.InterfaceBridge(d.Get("name").(string))
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}
		if _, ok := res["protocol-mode"]; ok {
			d.Set("protocol_mode", res["protocol-mode"])
		}
		if _, ok := res["priority"]; ok {
			d.Set("priority", res["priority"])
		}
		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikInterfaceBridgeUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetInterfaceBridgeComment(d.Get("name").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}
		if d.HasChange("protocol_mode") {
			if err := mt.SetInterfaceBridgeProtocolMode(d.Get("name").(string), d.Get("protocol_mode").(string)); err != nil {
				return err
			}
			d.SetPartial("protocol_mode")
		}
		if d.HasChange("priority") {
			if err := mt.SetInterfaceBridgePriority(d.Get("name").(string), d.Get("priority").(string)); err != nil {
				return err
			}
			d.SetPartial("priority")
		}

		// read-only parameters
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_interface_bridge.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikInterfaceBridgeDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
