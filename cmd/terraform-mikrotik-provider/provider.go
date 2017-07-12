package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/ozym/ros"
)

// Provider returns a schema.Provider for Example.
func mikrotikProvider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "admin",
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ssh_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  22,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "30s",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"mikrotik":                            resourceMikroTik(),
			"mikrotik_interface_ethernet":         resourceMikroTikInterfaceEthernet(),
			"mikrotik_interface_gre":              resourceMikroTikInterfaceGre(),
			"mikrotik_interface_bridge":           resourceMikroTikInterfaceBridge(),
			"mikrotik_interface_bridge_port":      resourceMikroTikInterfaceBridgePort(),
			"mikrotik_ip_address":                 resourceMikroTikIpAddress(),
			"mikrotik_ip_dns":                     resourceMikroTikIpDns(),
			"mikrotik_ip_route":                   resourceMikroTikIpRoute(),
			"mikrotik_ip_service":                 resourceMikroTikIpService(),
			"mikrotik_routing_filter_chain":       resourceMikroTikRoutingFilterChain(),
			"mikrotik_routing_bgp_instance":       resourceMikroTikRoutingBgpInstance(),
			"mikrotik_routing_bgp_aggregate":      resourceMikroTikRoutingBgpAggregate(),
			"mikrotik_routing_bgp_network":        resourceMikroTikRoutingBgpNetwork(),
			"mikrotik_routing_bgp_peer":           resourceMikroTikRoutingBgpPeer(),
			"mikrotik_routing_ospf_instance":      resourceMikroTikRoutingOspfInstance(),
			"mikrotik_routing_ospf_interface":     resourceMikroTikRoutingOspfInterface(),
			"mikrotik_routing_ospf_network":       resourceMikroTikRoutingOspfNetwork(),
			"mikrotik_routing_ospf_nbma_neighbor": resourceMikroTikRoutingOspfNbmaNeighbor(),
			"mikrotik_snmp":                       resourceMikroTikSnmp(),
			"mikrotik_system_clock":               resourceMikroTikSystemClock(),
			"mikrotik_system_identity":            resourceMikroTikSystemIdentity(),
			"mikrotik_system_logging":             resourceMikroTikSystemLogging(),
			"mikrotik_system_logging_action":      resourceMikroTikSystemLoggingAction(),
			"mikrotik_system_ntp_client":          resourceMikroTikSystemNtpClient(),
			"mikrotik_system_note":                resourceMikroTikSystemNote(),
			"mikrotik_system_script":              resourceMikroTikSystemScript(),
			"mikrotik_tool_netwatch":              resourceMikroTikToolNetwatch(),
			"mikrotik_tool_romon":                 resourceMikroTikToolRomon(),
			"mikrotik_tool_romon_port":            resourceMikroTikToolRomonPort(),
			"mikrotik_user":                       resourceMikroTikUser(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"mikrotik_system_resource":             dataSourceMikroTikSystemResource(),
			"mikrotik_interfaces":                  dataSourceMikroTikInterfaces(),
			"mikrotik_interface_bridge_ports":      dataSourceMikroTikInterfaceBridgePorts(),
			"mikrotik_interface_wirelesses":        dataSourceMikroTikInterfaceWirelesses(),
			"mikrotik_routing_ospf_instances":      dataSourceMikroTikRoutingOspfInstances(),
			"mikrotik_routing_ospf_interfaces":     dataSourceMikroTikRoutingOspfInterfaces(),
			"mikrotik_routing_ospf_networks":       dataSourceMikroTikRoutingOspfNetworks(),
			"mikrotik_routing_ospf_nbma_neighbors": dataSourceMikroTikRoutingOspfNbmaNeighbors(),
			"mikrotik_ip_addresses":                dataSourceMikroTikIpAddresses(),
			"mikrotik_version":                     dataSourceMikroTikVersion(),
		},

		ConfigureFunc: mikrotikConfigure,
	}
}

func mikrotikConfigure(d *schema.ResourceData) (interface{}, error) {
	hostname, port := d.Get("hostname").(string), d.Get("ssh_port").(int)
	username, password := d.Get("username").(string), d.Get("password").(string)

	timeout, err := time.ParseDuration(d.Get("timeout").(string))
	if err != nil {
		return nil, err
	}

	options := []func(*ros.Ros) error{
		ros.Port(port), ros.Username(username), ros.Password(password), ros.Timeout(timeout),
	}

	return ros.NewRos(hostname, options...)

}

func resourceRead(d *schema.ResourceData, meta interface{}, f func(d *schema.ResourceData, mt *ros.Ros) error) error {
	mt := meta.(*ros.Ros)

	d.SetId("")

	if mt.Error() != nil {
		return nil
	}

	if err := mt.Version(); err != nil {
		return nil
	}

	if err := f(d, mt); err != nil {
		return nil
	}

	return nil
}

func resourceWrite(d *schema.ResourceData, meta interface{}, f func(d *schema.ResourceData, mt *ros.Ros) error) error {
	mt := meta.(*ros.Ros)

	d.SetId("")

	if err := mt.Error(); err != nil {
		return err
	}

	if err := mt.Version(); err != nil {
		return err
	}

	return f(d, mt)
}

func resourceImport(d *schema.ResourceData, meta interface{}, f func(d *schema.ResourceData, mt *ros.Ros) ([]*schema.ResourceData, error)) ([]*schema.ResourceData, error) {
	mt := meta.(*ros.Ros)

	d.SetId("")

	if err := mt.Error(); err != nil {
		return nil, err
	}

	if err := mt.Version(); err != nil {
		return nil, err
	}

	return f(d, mt)
}
