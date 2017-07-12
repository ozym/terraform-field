package ros

func routingBgpPeers() Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingBgpPeers() ([]map[string]string, error) {
	return r.List(routingBgpPeers())
}

func routingBgpPeer(addr string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "print",
		Filter: map[string]string{
			"remote-address": addr,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBgpPeer(addr string) (map[string]string, error) {
	return r.First(routingBgpPeer(addr))
}

func setRoutingBgpPeer(addr, key, value string) Command {
	return Command{
		Path:    "/routing bgp peer",
		Command: "set",
		Filter: map[string]string{
			"remote-address": addr,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingBgpPeerComment(addr, comment string) error {
	return r.Exec(setRoutingBgpPeer(addr, "comment", comment))
}
func (r Ros) SetRoutingBgpPeerName(addr, name string) error {
	return r.Exec(setRoutingBgpPeer(addr, "name", name))
}
