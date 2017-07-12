package main

import (
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/ozym/skyedge"
)

func dataSourceSkyEdgeStatus() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSkyEdgeStatusRead,

		Schema: map[string]*schema.Schema{
			"vsat_pn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsat_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsat_sn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"outbound_freq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hw_mboard": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sw_bt_ver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sw_active_bt_ver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sw_op_ver": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"modcod": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSkyEdgeStatusRead(d *schema.ResourceData, meta interface{}) error {

	idu := meta.(*skyedge.SkyEdge)

	s, err := idu.Status()
	if err != nil {
		return err
	}

	//	 &{Operational 0  Up Synchronized Authenticated Full Access Up On Locked Off 100Mbps / Full duplex No connection Normal 047 days 10:21:48 Satellite: 0.00E Normal Inserted active / satellite link up 10007 00-A0-AC-18-95-3E 0 0 569004 10007 0411510492 0x4305 7.0.3.1 7.0.3.1 60.85.01.04 00-A0-AC-18-95-3E 202.139.38.29 255.255.255.252 172.27.149.62 255.255.0.0 9 13.88  1937800 DVB-S2 ACM 412375 456093}

	/*
	   StActiveSw  string
	   WebDivMode  string
	   StGwCon     string

	   StSatLink   string
	   StSync      string
	   StAuth      string
	   StAuthor    string
	   StNetLink   string
	   StTcpAcc    string
	   StObLock    string
	   StRpa       string
	*/

	/*
	   StLanPort   string
	   StLanPort2  string
	   StPwrMode   string
	*/

	/*
	   StOnTime    string
	   StActiveSat string
	   StMode      string
	   StSdCard    string
	   WebGwtStatus1 string
	*/
	/*

	   WebGwtId1     string
	   WebGwtMac1    string
	   WebGwrs       string
	   WebOthers     string
	*/

	/*
	   VsatPn string
	   VsatId string
	   VsatSn string
	*/

	/*
	   HwMBoard      string
	   SwBtVer       string
	   SwActiveBtVer string
	   SwOpVer       string
	*/

	/*
	   NetMac   string
	   NetIp    string
	   NetMask  string
	*/
	/*
	   NetAIp   string
	   NetAMask string

	   StCputil string

	   AccessMode      string
	   RxEbN0          string
	   OutboundFreq    string
	   ModCod          string
	   DrrpLanSent     string
	   DrrpLanReceived string

	   WebGwtStatus1 string
	   WebGwtId1     string
	   WebGwtMac1    string
	   WebGwrs       string
	   WebOthers     string

	   VsatPn string
	   VsatId string
	   VsatSn string

	   HwMBoard      string
	   SwBtVer       string
	   SwActiveBtVer string
	   SwOpVer       string

	   NetMac   string
	   NetIp    string
	   NetMask  string
	   NetAIp   string
	   NetAMask string

	   StCputil string

	   AccessMode      string
	   RxEbN0          string
	   OutboundFreq    string
	   ModCod          string
	   DrrpLanSent     string
	   DrrpLanReceived string
	*/

	d.Set("vsat_pn", s.VsatPn)
	d.Set("vsat_id", s.VsatId)
	d.Set("vsat_sn", s.VsatSn)
	d.Set("mac_addr", s.NetMac)
	d.Set("hw_mboard", s.HwMBoard)
	d.Set("sw_bt_ver", s.SwBtVer)
	d.Set("sw_active_bt_ver", s.SwActiveBtVer)
	d.Set("sw_op_ver", s.SwOpVer)

	if i, err := strconv.Atoi(s.OutboundFreq); err == nil {
		d.Set("outbound_freq", i)
	}
	d.Set("modcod", s.ModCod)

	d.SetId(idu.Id())

	return nil
}
