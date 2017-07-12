package main

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikIpService() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikIpServiceCreate,
		Read:   resourceMikroTikIpServiceRead,
		Update: resourceMikroTikIpServiceUpdate,
		Delete: resourceMikroTikIpServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"address": &schema.Schema{
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

func resourceMikroTikIpServiceCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikIpServiceUpdate(d, meta)
}

func resourceMikroTikIpServiceRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.IpService(d.Get("name").(string))
		if err != nil {
			return err
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		if _, ok := res["port"]; ok {
			p, err := strconv.Atoi(res["port"])
			if err != nil {
				return err
			}
			d.Set("port", p)
		}

		if _, ok := res["address"]; ok {
			d.Set("address", res["address"])
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikIpServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("disabled") {
			if d.Get("name").(string) != "ssh" {
				if err := mt.SetIpServiceDisabled(d.Get("name").(string), d.Get("disabled").(bool)); err != nil {
					return err
				}
			}
			d.SetPartial("disabled")
		}

		if d.HasChange("port") {
			if err := mt.SetIpServicePort(d.Get("name").(string), d.Get("port").(int)); err != nil {
				return err
			}
			d.SetPartial("port")
		}

		if d.HasChange("address") {
			if err := mt.SetIpServiceAddress(d.Get("name").(string), d.Get("address").(string)); err != nil {
				return err
			}
			d.SetPartial("address")
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikIpServiceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
