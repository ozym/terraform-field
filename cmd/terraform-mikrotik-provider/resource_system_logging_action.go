package main

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/ros"
)

func resourceMikroTikSystemLoggingAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceMikroTikSystemLoggingActionCreate,
		Read:   resourceMikroTikSystemLoggingActionRead,
		Update: resourceMikroTikSystemLoggingActionUpdate,
		Delete: resourceMikroTikSystemLoggingActionDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "remote",
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "remote",
			},
			"remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"remote_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"src_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"bsd_syslog": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"syslog_facility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "daemon",
			},
			"syslog_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "auto",
			},
		},
	}
}

func resourceMikroTikSystemLoggingActionCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if err := mt.SetSystemLoggingActionTarget(d.Get("name").(string), d.Get("target").(string)); err != nil {
			return err
		}
		if err := mt.SetSystemLoggingActionRemote(d.Get("name").(string), d.Get("remote").(string)); err != nil {
			return err
		}
		if err := mt.SetSystemLoggingActionRemotePort(d.Get("name").(string), d.Get("remote_port").(int)); err != nil {
			return err
		}
		if err := mt.SetSystemLoggingActionBsdSyslog(d.Get("name").(string), d.Get("bsd_syslog").(bool)); err != nil {
			return err
		}
		if err := mt.SetSystemLoggingActionSrcAddress(d.Get("name").(string), d.Get("src_address").(string)); err != nil {
			return err
		}
		if err := mt.SetSystemLoggingActionSyslogFacility(d.Get("name").(string), d.Get("syslog_facility").(string)); err != nil {
			return err
		}
		if err := mt.SetSystemLoggingActionSyslogSeverity(d.Get("name").(string), d.Get("syslog_severity").(string)); err != nil {
			return err
		}

		d.SetId(mt.Id() + ":::" + d.Get("name").(string))

		return resourceMikroTikSystemLoggingActionUpdate(d, meta)
	})
}

func resourceMikroTikSystemLoggingActionRead(d *schema.ResourceData, meta interface{}) error {
	return resourceRead(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if d.Get("name").(string) != "" {

			res, err := mt.SystemLoggingAction(d.Get("name").(string))
			if err != nil {
				return err
			}

			if _, ok := res["remote"]; ok {
				d.Set("remote", res["remote"])
			}
			if _, ok := res["target"]; ok {
				d.Set("target", res["target"])
			}
			if _, ok := res["bsd-syslog"]; ok {
				d.Set("bsd_syslog", ros.ParseBool(res["bsd-syslog"]))
			}
			if _, ok := res["syslog-severity"]; ok {
				d.Set("syslog_severity", res["syslog-severity"])
			}
			if _, ok := res["remote-port"]; ok {
				port, err := strconv.Atoi(res["remote-port"])
				if err != nil {
					return err
				}
				d.Set("remote_port", port)
			}
			if _, ok := res["src-address"]; ok {
				d.Set("src_address", res["src-address"])
			}
			if _, ok := res["syslog-facility"]; ok {
				d.Set("syslog_facility", res["syslog-facility"])
			}

			d.SetId(mt.Id() + ":::" + d.Get("name").(string))
		}

		return nil
	})
}

func resourceMikroTikSystemLoggingActionUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceWrite(d, meta, func(d *schema.ResourceData, mt *ros.Ros) error {

		if d.Get("name").(string) != "" {

			d.Partial(true)

			if d.HasChange("remote") {
				if err := mt.SetSystemLoggingActionRemote(d.Get("name").(string), d.Get("remote").(string)); err != nil {
					return err
				}
				d.SetPartial("remote")
			}

			if d.HasChange("target") {
				if err := mt.SetSystemLoggingActionTarget(d.Get("name").(string), d.Get("target").(string)); err != nil {
					return err
				}
				d.SetPartial("target")
			}

			if d.HasChange("bsd_syslog") {
				if err := mt.SetSystemLoggingActionBsdSyslog(d.Get("name").(string), d.Get("bsd_syslog").(bool)); err != nil {
					return err
				}
				d.SetPartial("bsd_syslog")
			}

			if d.HasChange("syslog_severity") {
				if err := mt.SetSystemLoggingActionSyslogSeverity(d.Get("name").(string), d.Get("syslog_severity").(string)); err != nil {
					return err
				}

				d.SetPartial("syslog_severity")
			}

			if d.HasChange("remote_port") {
				if err := mt.SetSystemLoggingActionRemotePort(d.Get("name").(string), d.Get("remote_port").(int)); err != nil {
					return err
				}
				d.SetPartial("remote_port")
			}

			if d.HasChange("src_address") {
				if err := mt.SetSystemLoggingActionSrcAddress(d.Get("name").(string), d.Get("src_address").(string)); err != nil {
					return err
				}
				d.SetPartial("src_address")
			}

			if d.HasChange("syslog_facility") {
				if err := mt.SetSystemLoggingActionSyslogFacility(d.Get("name").(string), d.Get("syslog_facility").(string)); err != nil {
					return err
				}
				d.SetPartial("syslog_facility")
			}

			d.Partial(false)

			d.SetId(mt.Id() + ":::" + d.Get("name").(string))
		}

		return nil
	})
}

func resourceMikroTikSystemLoggingActionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
