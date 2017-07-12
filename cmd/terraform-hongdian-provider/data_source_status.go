package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/hongdian"
)

func dataSourceHongdianStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHongdianStatusRead,

		Schema: map[string]*schema.Schema{
			"router_model": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardware_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"software_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"product_sn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceHongdianStatusRead(d *schema.ResourceData, meta interface{}) error {

	modem := meta.(*hongdian.Hongdian)

	s, err := modem.Status()
	if err != nil {
		return err
	}

	d.Set("router_model", s.RouterModel)
	d.Set("hardware_version", s.HardwareVersion)
	d.Set("software_version", s.SoftwareVersion)
	d.Set("serial_number", s.SerialNumber)
	d.Set("ip_address", s.IpAddress)
	d.Set("device_name", s.DeviceName)
	d.Set("vendor_id", s.VendorId)
	d.Set("product_id", s.ProductId)
	d.Set("product_sn", s.ProductSn)
	d.Set("network_type", s.NetworkType)

	d.SetId(modem.Id())

	return nil
}
