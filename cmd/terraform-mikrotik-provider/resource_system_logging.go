package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikSystemLogging() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikSystemLoggingCreate,
		Read:   resourceMikroTikSystemLoggingRead,
		Update: resourceMikroTikSystemLoggingUpdate,
		Delete: resourceMikroTikSystemLoggingDelete,

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"topics": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceMikroTikSystemLoggingCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemLogging(d.Get("action").(string), d.Get("topics").(string))
		if err != nil {
			return err
		}
		if res == nil {
			if err := mt.AddSystemLogging(d.Get("action").(string), d.Get("topics").(string)); err != nil {
				return err
			}
		}

		return resourceMikroTikSystemLoggingUpdate(d, meta)
	})
}

func resourceMikroTikSystemLoggingRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemLogging(d.Get("action").(string), d.Get("topics").(string))
		if err != nil {
			return err
		}
		if _, ok := res["prefix"]; ok {
			d.Set("prefix", res["prefix"])
		}

		d.SetId(mt.Id() + ":::" + d.Get("action").(string))

		return nil
	})
}

func resourceMikroTikSystemLoggingUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetSystemLoggingPrefix(d.Get("action").(string), d.Get("topics").(string), d.Get("prefix").(string)); err != nil {
			return err
		}

		return resourceMikroTikSystemLoggingRead(d, meta)
	})
}

func resourceMikroTikSystemLoggingDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.RemoveSystemLogging(d.Get("action").(string), d.Get("topics").(string)); err != nil {
			return err
		}
		return nil
	})
}
