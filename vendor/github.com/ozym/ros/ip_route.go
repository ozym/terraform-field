package ros

func ipRoutes(static bool) Command {
	return Command{
		Path:    "/ip route",
		Command: "print",
		Flags: map[string]bool{
			"static": static,
		},
		Detail: true,
	}
}

func (r Ros) IpRoutes(static bool) ([]map[string]string, error) {
	return r.List(ipRoutes(static))
}

func ipRoute(address string) Command {
	return Command{
		Path:    "/ip route",
		Command: "print",
		Filter: map[string]string{
			"dst-address": address,
		},
		Detail: true,
	}
}

func (r Ros) IpRoute(address string) (map[string]string, error) {
	return r.First(ipAddress(address))
}

func setIpRoute(address, key, value string) Command {
	return Command{
		Path:    "/ip route",
		Command: "set",
		Filter: map[string]string{
			"dst-address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetIpRouteComment(address, comment string) error {
	return r.Exec(setIpRoute(address, "comment", comment))
}
