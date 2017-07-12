package ros

func routingBgpInstances() Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingBgpInstances() ([]map[string]string, error) {
	return r.List(routingBgpInstances())
}

func routingBgpInstance(name string) Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBgpInstance(name string) (map[string]string, error) {
	return r.First(routingBgpInstance(name))
}

func setRoutingBgpInstance(name, key, value string) Command {
	return Command{
		Path:    "/routing bgp instance",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetRoutingBgpInstanceRouterId(name, router_id string) error {
	return r.Exec(setRoutingBgpInstance(name, "router-id", router_id))
}

func (r Ros) SetRoutingBgpInstanceComment(name, comment string) error {
	return r.Exec(setRoutingBgpInstance(name, "comment", comment))
}
