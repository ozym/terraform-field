package ros

func routingFilters() Command {
	return Command{
		Path:    "/routing filter",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) RoutingFilters() ([]map[string]string, error) {
	return r.List(routingFilters())
}

func routingFilterChain(chain string) Command {
	return Command{
		Path:    "/routing filter",
		Command: "print",
		Filter: map[string]string{
			"chain": chain,
		},

		Detail: true,
	}
}

func (r Ros) RoutingFilterChain(chain string) ([]map[string]string, error) {
	return r.List(routingFilterChain(chain))
}
