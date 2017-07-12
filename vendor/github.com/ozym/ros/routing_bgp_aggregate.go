package ros

func routingBgpAggregates() Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingBgpAggregates() ([]map[string]string, error) {
	return r.List(routingBgpAggregates())
}

func routingBgpAggregate(instance, prefix string) Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "print",
		Filter: map[string]string{
			"instance": instance,
			"prefix":   prefix,
		},
		Detail: true,
	}
}

func (r Ros) RoutingBgpAggregate(instance, prefix string) (map[string]string, error) {
	return r.First(routingBgpAggregate(instance, prefix))
}

func setRoutingBgpAggregate(instance, prefix, key, value string) Command {
	return Command{
		Path:    "/routing bgp aggregate",
		Command: "set",
		Filter: map[string]string{
			"instance": instance,
			"prefix":   prefix,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetRoutingBgpAggregateComment(instance, prefix, comment string) error {
	return r.Exec(setRoutingBgpAggregate(instance, prefix, "comment", comment))
}
