package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikIpDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikIpDnsCreate,
		Read:   resourceMikroTikIpDnsRead,
		Update: resourceMikroTikIpDnsUpdate,
		Delete: resourceMikroTikIpDnsDelete,

		Schema: map[string]*schema.Schema{
			"servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"allow_remote_requests": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceMikroTikIpDnsCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetIpDnsServers(d.Get("servers").(string)); err != nil {
			return err
		}

		if err := mt.SetIpDnsAllowRemoteRequests(d.Get("allow_remote_requests").(bool)); err != nil {
			return err
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikIpDnsRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.IpDns()
		if err != nil {
			return err
		}

		if _, ok := res["servers"]; ok {
			d.Set("servers", res["servers"])
		}
		if _, ok := res["allow-remote-requests"]; ok {
			d.Set("allow_remote_requests", ros.ParseBool(res["allow-remote-requests"]))
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikIpDnsUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("servers") {
			if err := mt.SetIpDnsServers(d.Get("servers").(string)); err != nil {
				return err
			}
			d.SetPartial("servers")
		}

		if d.HasChange("allow_remote_requests") {
			if err := mt.SetIpDnsAllowRemoteRequests(d.Get("allow_remote_requests").(bool)); err != nil {
				return err
			}
			d.SetPartial("allow_remote_requests")
		}

		d.Partial(false)

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikIpDnsDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetIpDnsServers(""); err != nil {
			return err
		}

		if err := mt.SetIpDnsAllowRemoteRequests(false); err != nil {
			return err
		}

		return nil
	})
}
