package main

import (
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/airos"
)

func dataSourceUbiquitiRadioStations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUbiquitiRadioStationsRead,

		Schema: map[string]*schema.Schema{
			"stations": &schema.Schema{
				Type:     schema.TypeSet,
				ForceNew: true,
				Computed: true,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"mac_addr": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"last_ip": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"tx_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"rx_rate": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
				Set: func(v interface{}) int {
					m := v.(map[string]interface{})
					return hashcode.String(m["name"].(string))
				},
			},
		},
	}
}

func dataSourceUbiquitiRadioStationsRead(d *schema.ResourceData, meta interface{}) error {

	radio := meta.(*airos.AirOs)

	stns, err := radio.Stations()
	if err != nil {
		return err
	}

	stations := []map[string]interface{}{}

	for _, s := range stns {
		in := map[string]interface{}{}

		in["name"] = s.Name
		in["mac_addr"] = s.MacAddr
		in["last_ip"] = s.LastIp
		in["tx_rate"] = s.TxRate
		in["rx_rate"] = s.RxRate

		stations = append(stations, in)
	}

	if err := d.Set("stations", stations); err != nil {
		return err
	}

	d.SetId(radio.Id())

	return nil
}
