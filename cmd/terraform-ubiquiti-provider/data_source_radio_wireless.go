package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/airos"
)

func dataSourceUbiquitiRadioWireless() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUbiquitiRadioWirelessRead,

		Schema: map[string]*schema.Schema{
			"ssid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hide_ssid": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ap_mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tx_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"rx_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"wds_enabled": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"channel_width": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"station_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceUbiquitiRadioWirelessRead(d *schema.ResourceData, meta interface{}) error {

	radio := meta.(*airos.AirOs)

	err := radio.Connect()
	if err != nil {
		return err
	}

	w, err := radio.Wireless()
	if err != nil {
		return err
	}

	d.Set("ssid", w.Ssid)
	d.Set("hide_ssid", w.HideSsid)
	d.Set("ap_mac_addr", w.ApMacAddr)
	d.Set("tx_rate", w.TxRate)
	d.Set("rx_rate", w.RxRate)
	d.Set("security", w.Security)
	d.Set("wds_enabled", w.WdsEnabled)
	d.Set("channel_width", w.ChannelWidth)
	d.Set("station_count", w.StationCount)

	d.SetId(radio.Id())

	return nil
}
