package ros

func routingOspfInterfaces() Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingOspfInterfaces() ([]map[string]string, error) {
	return r.List(routingOspfInterfaces())
}

func routingOspfInterface(iface string) Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "print",
		Filter: map[string]string{
			"interface": iface,
		},
		Detail: true,
	}
}

func (r Ros) RoutingOspfInterface(iface string) (map[string]string, error) {
	return r.First(routingOspfInterface(iface))
}

func setRoutingOspfInterface(iface, key, value string) Command {
	return Command{
		Path:    "/routing ospf interface",
		Command: "set",
		Filter: map[string]string{
			"interface": iface,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingOspfInterfaceComment(iface, comment string) error {
	return r.Exec(setRoutingOspfInterface(iface, "comment", comment))
}
func (r Ros) SetRoutingOspfInterfaceCost(iface, cost string) error {
	return r.Exec(setRoutingOspfInterface(iface, "cost", cost))
}
