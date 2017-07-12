package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikSystemNote() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikSystemNoteCreate,
		Read:   resourceMikroTikSystemNoteRead,
		Update: resourceMikroTikSystemNoteUpdate,
		Delete: resourceMikroTikSystemNoteDelete,

		Schema: map[string]*schema.Schema{
			"note": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"show_at_login": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceMikroTikSystemNoteCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikSystemNoteUpdate(d, meta)
}

func resourceMikroTikSystemNoteRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemNote()
		if err != nil {
			return err
		}

		if _, ok := res["note"]; ok {
			d.Set("note", res["note"])
		}
		if _, ok := res["show-at-login"]; ok {
			d.Set("show_at_login", ros.ParseBool(res["show-at-login"]))
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemNoteUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("note") {
			if err := mt.SetSystemNote(d.Get("note").(string)); err != nil {
				return err
			}
			d.SetPartial("note")
		}

		if d.HasChange("show_at_login") {
			if err := mt.SetSystemNoteShowAtLogin(d.Get("show_at_login").(bool)); err != nil {
				return err
			}
			d.SetPartial("show_at_login")
		}

		d.Partial(false)

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemNoteDelete(d *schema.ResourceData, meta interface{}) error {

	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetSystemNote(""); err != nil {
			return err
		}

		if err := mt.SetSystemNoteShowAtLogin(false); err != nil {
			return err
		}

		return nil
	})
}
