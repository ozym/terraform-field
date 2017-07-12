package ros

func routingBgpNetworks() Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingBgpNetworks() ([]map[string]string, error) {
	return r.List(routingBgpNetworks())
}

func routingBgpNetwork(network string) Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "print",
		Filter: map[string]string{
			"network": network,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBgpNetwork(network string) (map[string]string, error) {
	return r.First(routingBgpNetwork(network))
}

func setRoutingBgpNetwork(network, key, value string) Command {
	return Command{
		Path:    "/routing bgp network",
		Command: "set",
		Filter: map[string]string{
			"network": network,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingBgpNetworkComment(network, comment string) error {
	return r.Exec(setRoutingBgpNetwork(network, "comment", comment))
}
