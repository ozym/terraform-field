package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikSystemClock() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikSystemClockCreate,
		Read:   resourceMikroTikSystemClockRead,
		Update: resourceMikroTikSystemClockUpdate,
		Delete: resourceMikroTikSystemClockDelete,

		Schema: map[string]*schema.Schema{
			"time_zone_autodetect": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"time_zone_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Pacific/Auckland",
			},
		},
	}
}

func resourceMikroTikSystemClockHasTimeZoneAutodetect(d *schema.ResourceData, mt *ros.Ros) bool {

	switch {
	case mt.Major() < 6:
		return false
	case mt.Major() > 6:
		return true
	case mt.Minor() < 27:
		return false
	default:
		return true
	}
}

func resourceMikroTikSystemClockCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetSystemClockTimeZoneName(d.Get("time_zone_name").(string)); err != nil {
			return err
		}
		if resourceMikroTikSystemClockHasTimeZoneAutodetect(d, mt) {
			if err := mt.SetSystemClockTimeZoneAutodetect(d.Get("time_zone_autodetect").(bool)); err != nil {
				return err
			}
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemClockRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.SystemClock()
		if err != nil {
			return err
		}

		if _, ok := res["time-zone-name"]; ok {
			d.Set("time_zone_name", res["time-zone-name"])
		}

		if resourceMikroTikSystemClockHasTimeZoneAutodetect(d, mt) {
			if _, ok := res["time-zone-autodetect"]; ok {
				d.Set("time_zone_autodetect", ros.ParseBool(res["time-zone-autodetect"]))
			}
		} else {
			d.Set("time_zone_autodetect", false)
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemClockUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		d.Partial(true)

		if d.HasChange("time_zone_name") {
			if err := mt.SetSystemClockTimeZoneName(d.Get("time_zone_name").(string)); err != nil {
				return err
			}
			d.SetPartial("time_zone_name")
		}

		if d.HasChange("time_zone_autodetect") {
			if resourceMikroTikSystemClockHasTimeZoneAutodetect(d, mt) {
				if err := mt.SetSystemClockTimeZoneAutodetect(d.Get("time_zone_autodetect").(bool)); err != nil {
					return err
				}
			}
			d.SetPartial("time_zone_autodetect")
		}

		d.Partial(false)

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikSystemClockDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
