package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/airos"
)

func dataSourceUbiquitiRadioStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUbiquitiRadioStatusRead,

		Schema: map[string]*schema.Schema{
			"tx_power": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"radio_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"country_code": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"dfs_enabled": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"antenna": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceUbiquitiRadioStatusRead(d *schema.ResourceData, meta interface{}) error {

	radio := meta.(*airos.AirOs)

	err := radio.Connect()
	if err != nil {
		return err
	}

	s, err := radio.Status()
	if err != nil {
		return err
	}

	d.Set("radio_mode", s.RadioMode)
	d.Set("country_code", s.CountryCode)
	d.Set("frequency", s.Frequency)
	d.Set("dfs_enabled", s.DfsEnabled)
	d.Set("tx_power", s.TxPower)
	d.Set("distance", s.Distance)
	d.Set("antenna", s.Antenna)

	d.SetId(radio.Id())

	return nil
}
