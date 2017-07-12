package main

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ozym/ros"
)

func resourceMikroTikInterfaceEthernet() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikInterfaceEthernetCreate,
		Read:   resourceMikroTikInterfaceEthernetRead,
		Update: resourceMikroTikInterfaceEthernetUpdate,
		Delete: resourceMikroTikInterfaceEthernetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceMikroTikInterfaceEthernetImporter,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
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

func resourceMikroTikInterfaceEthernetCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMikroTikInterfaceEthernetUpdate(d, meta)
}

func resourceMikroTikInterfaceEthernetRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		res, err := mt.InterfaceEthernet(d.Get("name").(string))
		if err != nil {
			return err
		}

		if _, ok := res["comment"]; ok {
			d.Set("comment", res["comment"])
		}

		if _, ok := res["disabled"]; ok {
			d.Set("disabled", ros.ParseBool(res["disabled"]))
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikInterfaceEthernetUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {
		d.Partial(true)

		if d.HasChange("comment") {
			if err := mt.SetInterfaceEthernetComment(d.Get("name").(string), d.Get("comment").(string)); err != nil {
				return err
			}
			d.SetPartial("comment")
		}

		// read-only
		if d.HasChange("disabled") {
			return fmt.Errorf("unable to set read-only mikrotik_interface_ethernet.disabled parameter on %s", d.Get("name").(string))
		}

		d.Partial(false)

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return nil
	})
}

func resourceMikroTikInterfaceEthernetDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceMikroTikInterfaceEthernetImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return resourceImport(d, meta, func(d *schema.ResourceData, mt *ros.Ros) ([]*schema.ResourceData, error) {

		list, err := mt.InterfaceEthernets()
		if err != nil {
			return nil, err
		}

		var results []*schema.ResourceData

		for _, iface := range list {
			if name, ok := iface["name"]; ok {

				ethernet := resourceMikroTikInterfaceEthernet()
				eData := ethernet.Data(nil)
				eData.SetId(mt.Id() + ":::" + name)
				eData.SetType("mikrotik_interface_ethernet")
				eData.Set("name", name)
				if comment, ok := iface["comment"]; ok {
					eData.Set("comment", comment)
				}
				if disabled, ok := iface["disabled"]; ok {
					eData.Set("disabled", ros.ParseBool(disabled))
				}
				results = append(results, eData)
			}
		}

		return results, nil
	})
}
