package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikUserCreate,
		Read:   resourceMikroTikUserRead,
		Update: resourceMikroTikUserUpdate,
		Delete: resourceMikroTikUserDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceMikroTikUserCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		name := d.Get("name").(string)
		group := d.Get("group").(string)
		password := d.Get("password").(string)

		if name != "admin" {
			if err := mt.AddUser(name, group, password); err != nil {
				return err
			}
		}

		d.SetId(mt.Id() + ":::" + name)

		return nil
	})
}

func resourceMikroTikUserRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		name := d.Get("name").(string)

		res, err := mt.User(name)
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}
		if _, ok := res["group"]; ok {
			d.Set("group", res["group"])
		}

		d.SetId(func() string {
			if len(res) > 0 {
				return mt.Id() + ":::" + name
			}
			return ""
		}())

		return nil
	})
}

func resourceMikroTikUserUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		name := d.Get("name").(string)
		comment := d.Get("comment").(string)
		group := d.Get("group").(string)

		if d.HasChange("comment") {
			if err := mt.SetUserComment(name, comment); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		if d.HasChange("group") {
			if err := mt.SetUserGroup(name, group); err != nil {
				return err
			}
			d.SetPartial("group")
		}

		d.Partial(false)
		d.SetId(mt.Id() + ":::" + name)

		return nil
	})
}

func resourceMikroTikUserDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		name := d.Get("name").(string)

		if name != "admin" {
			if err := mt.RemoveUser(name); err != nil {
				return err
			}
		}

		return nil
	})
}
