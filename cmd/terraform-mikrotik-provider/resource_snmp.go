package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikSnmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikSnmpCreate,
		Read:   resourceMikroTikSnmpRead,
		Update: resourceMikroTikSnmpUpdate,
		Delete: resourceMikroTikSnmpDelete,

		Schema: map[string]*schema.Schema{
			"contact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceMikroTikSnmpCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetSnmpEnabled(d.Get("enabled").(bool)); err != nil {
			return err
		}
		if err := mt.SetSnmpContact(d.Get("contact").(string)); err != nil {
			return err
		}
		if err := mt.SetSnmpLocation(d.Get("location").(string)); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSnmpRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.Snmp()
		if err != nil {
			return err
		}

		if _, ok := res["enabled"]; ok {
			d.Set("enabled", ros.ParseBool(res["enabled"]))
		}
		if _, ok := res["contact"]; ok {
			d.Set("contact", res["contact"])
		}
		if _, ok := res["location"]; ok {
			d.Set("location", res["location"])
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSnmpUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("enabled") {
			if err := mt.SetSnmpEnabled(d.Get("enabled").(bool)); err != nil {
				return err
			}
			d.SetPartial("enabled")
		}

		if d.HasChange("contact") {
			if err := mt.SetSnmpContact(d.Get("contact").(string)); err != nil {
				return err
			}
			d.SetPartial("contact")
		}

		if d.HasChange("location") {
			if err := mt.SetSnmpLocation(d.Get("location").(string)); err != nil {
				return err
			}
			d.SetPartial("location")
		}

		d.Partial(false)

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSnmpDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		if err := mt.SetSnmpEnabled(false); err != nil {
			return err
		}
		return nil
	})
}
