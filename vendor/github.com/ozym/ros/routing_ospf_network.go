package ros

func routingOspfNetworks() Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingOspfNetworks() ([]map[string]string, error) {
	return r.List(routingOspfNetworks())
}

func routingOspfNetwork(network string) Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "print",
		Filter: map[string]string{
			"network": network,
		},
		Detail: true,
	}
}

func (r Ros) RoutingOspfNetwork(network string) (map[string]string, error) {
	return r.First(routingOspfNetwork(network))
}

func setRoutingOspfNetwork(network, key, value string) Command {
	return Command{
		Path:    "/routing ospf network",
		Command: "set",
		Filter: map[string]string{
			"network": network,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingOspfNetworkComment(network, comment string) error {
	return r.Exec(setRoutingOspfNetwork(network, "comment", comment))
}
