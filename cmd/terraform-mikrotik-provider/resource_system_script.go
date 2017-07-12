package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikSystemScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikSystemScriptCreate,
		Read:   resourceMikroTikSystemScriptRead,
		Update: resourceMikroTikSystemScriptUpdate,
		Delete: resourceMikroTikSystemScriptDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policy": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMikroTikSystemScriptCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemScript(d.Get("name").(string))
		if err != nil {
			return err
		}
		if res == nil {
			if err := mt.AddSystemScript(d.Get("name").(string), d.Get("policy").(string), d.Get("source").(string)); err != nil {
				return err
			}
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return resourceMikroTikSystemScriptUpdate(d, meta)
	})
}

func resourceMikroTikSystemScriptRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemScript(d.Get("name").(string))
		if err != nil {
			return err
		}
		if _, ok := res["policy"]; ok {
			d.Set("policy", res["policy"])
		}
		if _, ok := res["source"]; ok {
			d.Set("source", res["source"])
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikSystemScriptUpdate(d *schema.ResourceData, meta interface{}) error {

	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("policy") {
			if err := mt.SetSystemScriptPolicy(d.Get("name").(string), d.Get("policy").(string)); err != nil {
				return err
			}
			d.SetPartial("policy")
		}
		if d.HasChange("source") {
			if err := mt.SetSystemScriptSource(d.Get("name").(string), d.Get("source").(string)); err != nil {
				return err
			}
			d.SetPartial("source")
		}
		d.Partial(false)

		return resourceMikroTikSystemScriptRead(d, meta)
	})
}

func resourceMikroTikSystemScriptDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.RemoveSystemScript(d.Get("name").(string)); err != nil {
			return err
		}

		return nil
	})
}
