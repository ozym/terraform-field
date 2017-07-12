package main

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikToolRomonPort() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikToolRomonPortCreate,
		Read:   resourceMikroTikToolRomonPortRead,
		Update: resourceMikroTikToolRomonPortUpdate,
		Delete: resourceMikroTikToolRomonPortDelete,

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"forbid": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"cost": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"secrets": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceMikroTikToolRomonPortCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if resourceMikroTikToolHasRomon(d, mt) {
			legacy := resourceMikroTikToolHasLegacyRomon(d, mt)
			if d.Get("interface").(string) != "default" {
				if err := mt.AddToolRomonPort(d.Get("interface").(string), legacy); err != nil {
					return err
				}
			}
		}

		d.SetId(mt.Id() + ":::" + d.Get("interface").(string))

		return nil
	})
}

func resourceMikroTikToolRomonPortRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		if resourceMikroTikToolHasRomon(d, mt) {
			legacy := resourceMikroTikToolHasLegacyRomon(d, mt)

			res, err := mt.ToolRomonPort(d.Get("interface").(string), legacy)
			if err != nil {
				return err
			}

			if _, ok := res["disabled"]; ok {
				d.Set("disabled", ros.ParseBool(res["disabled"]))
			}
			if _, ok := res["forbid"]; ok {
				d.Set("forbid", ros.ParseBool(res["forbid"]))
			}
			if _, ok := res["cost"]; ok {
				c, err := strconv.Atoi(res["cost"])
				if err != nil {
					return err
				}
				d.Set("cost", c)
			}
			if _, ok := res["secrets"]; ok {
				d.Set("secrets", res["secrets"])
			}
		}

		d.SetId(mt.Id() + ":::" + d.Get("interface").(string))

		return nil
	})
}

func resourceMikroTikToolRomonPortUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		if resourceMikroTikToolHasRomon(d, mt) {
			legacy := resourceMikroTikToolHasLegacyRomon(d, mt)

			d.Partial(true)

			if d.HasChange("disabled") {
				if err := mt.SetToolRomonPortDisabled(d.Get("interface").(string), d.Get("disabled").(bool), legacy); err != nil {
					return err
				}
				d.SetPartial("disabled")
			}

			if d.HasChange("forbid") {
				if err := mt.SetToolRomonPortForbid(d.Get("interface").(string), d.Get("forbid").(bool), legacy); err != nil {
					return err
				}
				d.SetPartial("forbid")
			}

			if d.HasChange("cost") {
				if err := mt.SetToolRomonPortCost(d.Get("interface").(string), d.Get("cost").(int), legacy); err != nil {
					return err
				}
				d.SetPartial("cos")
			}

			if d.HasChange("secrets") {
				if err := mt.SetToolRomonPortSecrets(d.Get("interface").(string), d.Get("secrets").(string), legacy); err != nil {
					return err
				}
				d.SetPartial("secrets")
			}

			d.Partial(false)
		}

		d.SetId(mt.Id() + ":::" + d.Get("interface").(string))

		return nil
	})
}

func resourceMikroTikToolRomonPortDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		if resourceMikroTikToolHasRomon(d, mt) {
			legacy := resourceMikroTikToolHasLegacyRomon(d, mt)
			if d.Get("interface").(string) != "default" {
				if err := mt.RemoveToolRomonPort(d.Get("interface").(string), legacy); err != nil {
					return err
				}
			}
		}

		return nil
	})
}
