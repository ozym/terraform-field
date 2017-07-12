package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/fgr"
)

func dataSourceFreewaveRadioStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFreewaveRadioStatusRead,

		Schema: map[string]*schema.Schema{
			"model": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"firmware": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"contact": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"connected_to": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFreewaveRadioStatusRead(d *schema.ResourceData, meta interface{}) error {

	radio := meta.(*fgr.Fgr)

	err := radio.Connect()
	if err != nil {
		return err
	}

	s, err := radio.Status()
	if err != nil {
		return err
	}

	d.Set("model", s.Model)
	d.Set("version", s.Version)
	d.Set("firmware", s.Firmware)
	d.Set("contact", s.Contact)
	d.Set("name", s.Name)
	d.Set("location", s.Location)
	d.Set("connected_to", s.ConnectedTo)

	d.SetId(radio.Id())

	return nil
}
