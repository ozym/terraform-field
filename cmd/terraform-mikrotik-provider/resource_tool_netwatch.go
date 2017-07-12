package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikToolNetwatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikToolNetwatchCreate,
		Read:   resourceMikroTikToolNetwatchRead,
		Update: resourceMikroTikToolNetwatchUpdate,
		Delete: resourceMikroTikToolNetwatchDelete,

		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"up_script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"down_script": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "998ms",
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceMikroTikToolNetwatchCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.ToolNetwatch(d.Get("host").(string))
		if err != nil {
			return err
		}
		if res == nil {
			if err := mt.AddToolNetwatch(d.Get("host").(string), d.Get("up_script").(string), d.Get("down_script").(string), d.Get("interval").(string), d.Get("timeout").(string), d.Get("disabled").(bool)); err != nil {
				return err
			}
		}

		d.SetId(mt.Id() + ":::" + d.Get("host").(string))

		return resourceMikroTikToolNetwatchUpdate(d, meta)
	})
}

func resourceMikroTikToolNetwatchRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.ToolNetwatch(d.Get("host").(string))
		if err != nil {
			return err
		}
		if _, ok := res["up-script"]; ok {
			d.Set("up_script", res["up-script"])
		}
		if _, ok := res["down-script"]; ok {
			d.Set("down_script", res["down-script"])
		}
		if _, ok := res["interval"]; ok {
			d.Set("interval", res["interval"])
		}
		if _, ok := res["disabled"]; ok {
			d.Set("disabled", res["disabled"])
		}
		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		d.SetId(mt.Id() + ":::" + d.Get("host").(string))

		return nil
	})
}

func resourceMikroTikToolNetwatchUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("up_script") {
			if err := mt.SetToolNetwatchUpScript(d.Get("host").(string), d.Get("up_script").(string)); err != nil {
				return err
			}
			d.SetPartial("up_script")
		}
		if d.HasChange("down_script") {
			if err := mt.SetToolNetwatchDownScript(d.Get("host").(string), d.Get("down_script").(string)); err != nil {
				return err
			}
			d.SetPartial("down_script")
		}
		if d.HasChange("interval") {
			if err := mt.SetToolNetwatchInterval(d.Get("host").(string), d.Get("interval").(string)); err != nil {
				return err
			}
			d.SetPartial("interval")
		}
		if d.HasChange("timeout") {
			if err := mt.SetToolNetwatchTimeout(d.Get("host").(string), d.Get("timeout").(string)); err != nil {
				return err
			}
			d.SetPartial("timeout")
		}
		if d.HasChange("disabled") {
			if err := mt.SetToolNetwatchDisabled(d.Get("host").(string), d.Get("disabled").(bool)); err != nil {
				return err
			}
			d.SetPartial("disabled")
		}
		if d.HasChange("comment") {
			if err := mt.SetToolNetwatchComment(d.Get("host").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}
		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("host").(string))

		return nil
	})
}

func resourceMikroTikToolNetwatchDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.RemoveToolNetwatch(d.Get("host").(string)); err != nil {
			return err
		}

		return nil
	})
}
