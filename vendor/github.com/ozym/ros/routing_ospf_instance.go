package ros

func routingOspfInstances() Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingOspfInstances() ([]map[string]string, error) {
	return r.List(routingOspfInstances())
}

func routingOspfInstance(name string) Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) RoutingOspfInstance(name string) (map[string]string, error) {
	return r.First(routingOspfInstance(name))
}

func setRoutingOspfInstance(name, key, value string) Command {
	return Command{
		Path:    "/routing ospf instance",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingOspfInstanceRouterId(name, router_id string) error {
	return r.Exec(setRoutingOspfInstance(name, "router-id", router_id))
}
func (r Ros) SetRoutingOspfInstanceComment(name, comment string) error {
	return r.Exec(setRoutingOspfInstance(name, "comment", comment))
}
