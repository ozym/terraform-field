package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikToolRomon() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikToolRomonCreate,
		Read:   resourceMikroTikToolRomonRead,
		Update: resourceMikroTikToolRomonUpdate,
		Delete: resourceMikroTikToolRomonDelete,

		Schema: map[string]*schema.Schema{
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"mac_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "00:00:00:00:00:00",
			},
			"secrets": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"current_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceMikroTikToolHasRomon(d *schema.ResourceData, mt *ros.Ros) bool {

	switch {
	case mt.Major() < 6:
		return false
	case mt.Major() > 6:
		return true
	case mt.Minor() < 28:
		return false
	default:
		return true
	}
}

func resourceMikroTikToolHasLegacyRomon(d *schema.ResourceData, mt *ros.Ros) bool {

	switch {
	case mt.Major() < 6:
		return false
	case mt.Major() > 6:
		return false
	case mt.Minor() != 28:
		return false
	default:
		return true
	}
}

func resourceMikroTikToolRomonCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikToolRomonUpdate(d, meta)
}

func resourceMikroTikToolRomonRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if resourceMikroTikToolHasRomon(d, mt) {
			legacy := resourceMikroTikToolHasLegacyRomon(d, mt)

			res, err := mt.ToolRomon(legacy)
			if err != nil {
				return err
			}

			if _, ok := res["enabled"]; ok {
				d.Set("enabled", ros.ParseBool(res["enabled"]))
			}
			if _, ok := res["id"]; ok {
				d.Set("mac_id", res["id"])
			}
			if _, ok := res["secrets"]; ok {
				d.Set("secrets", res["secrets"])
			}
			if _, ok := res["current-id"]; ok {
				d.Set("current_id", res["current-id"])
			}
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikToolRomonUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		if resourceMikroTikToolHasRomon(d, mt) {
			legacy := resourceMikroTikToolHasLegacyRomon(d, mt)

			res, err := mt.ToolRomon(legacy)
			if err != nil {
				return err
			}

			d.Partial(true)

			if d.HasChange("enabled") {
				if err := mt.SetToolRomonEnabled(d.Get("enabled").(bool), legacy); err != nil {
					return err
				}
				d.SetPartial("enabled")
			}

			if d.HasChange("mac_id") {
				if err := mt.SetToolRomonId(d.Get("mac_id").(string), legacy); err != nil {
					return err
				}
				d.SetPartial("mac_id")
			}

			if d.HasChange("secrets") {
				if err := mt.SetToolRomonSecrets(d.Get("secrets").(string), legacy); err != nil {
					return err
				}
				d.SetPartial("secrets")
			}

			d.Partial(false)

			if _, ok := res["current-id"]; ok {
				d.Set("current_id", res["current-id"])
			}
		}

		d.SetId(mt.Id())

		return nil
	})
}

func resourceMikroTikToolRomonDelete(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if resourceMikroTikToolHasRomon(d, mt) {
			legacy := resourceMikroTikToolHasLegacyRomon(d, mt)

			if err := mt.SetToolRomonEnabled(false, legacy); err != nil {
				return err
			}
			if err := mt.SetToolRomonId("00:00:00:00:00:00", legacy); err != nil {
				return err
			}
			if err := mt.SetToolRomonSecrets("", legacy); err != nil {
				return err
			}
		}

		return nil
	})
}
