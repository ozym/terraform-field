package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikSystemNtpClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikSystemNtpClientCreate,
		Read:   resourceMikroTikSystemNtpClientRead,
		Update: resourceMikroTikSystemNtpClientUpdate,
		Delete: resourceMikroTikSystemNtpClientDelete,

		Schema: map[string]*schema.Schema{
			"primary_ntp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secondary_ntp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceMikroTikSystemNtpClientCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetSystemNtpClientPrimaryNtp(d.Get("primary_ntp").(string)); err != nil {
			return err
		}
		if err := mt.SetSystemNtpClientSecondaryNtp(d.Get("secondary_ntp").(string)); err != nil {
			return err
		}
		if err := mt.SetSystemNtpClientEnabled(d.Get("enabled").(bool)); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemNtpClientRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemNtpClient()
		if err != nil {
			return err
		}

		if _, ok := res["primary-ntp"]; ok {
			d.Set("primary_ntp", res["primary-ntp"])
		}
		if _, ok := res["secondary-ntp"]; ok {
			d.Set("secondary_ntp", res["secondary-ntp"])
		}
		if _, ok := res["enabled"]; ok {
			d.Set("enabled", ros.ParseBool(res["enabled"]))
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemNtpClientUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("enabled") {
			if err := mt.SetSystemNtpClientEnabled(d.Get("enabled").(bool)); err != nil {
				return err
			}
			d.SetPartial("enabled")
		}

		if d.HasChange("primary_ntp") {
			if err := mt.SetSystemNtpClientPrimaryNtp(d.Get("primary_ntp").(string)); err != nil {
				return err
			}
			d.SetPartial("primary_ntp")
		}

		if d.HasChange("secondary_ntp") {
			if err := mt.SetSystemNtpClientSecondaryNtp(d.Get("secondary_ntp").(string)); err != nil {
				return err
			}
			d.SetPartial("secondary_ntp")
		}

		d.Partial(false)

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemNtpClientDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetSystemNtpClientEnabled(false); err != nil {
			return err
		}
		if err := mt.SetSystemNtpClientPrimaryNtp(""); err != nil {
			return err
		}
		if err := mt.SetSystemNtpClientSecondaryNtp(""); err != nil {
			return err
		}

		return nil
	})
}
