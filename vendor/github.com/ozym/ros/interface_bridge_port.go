package ros

func interfaceBridgePorts() Command {
	return Command{
		Path:    "/interface bridge port",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) InterfaceBridgePorts() ([]map[string]string, error) {
	return r.List(interfaceBridgePorts())
}

func interfaceBridgePort(bridge, iface string) Command {
	return Command{
		Path:    "/interface bridge port",
		Command: "print",
		Filter: map[string]string{
			"bridge":    bridge,
			"interface": iface,
		},
		Detail: true,
	}
}

func (r Ros) InterfaceBridgePort(bridge, iface string) (map[string]string, error) {
	return r.First(interfaceBridgePort(bridge, iface))
}

func setInterfaceBridgePort(bridge, iface, key, value string) Command {
	return Command{
		Path:    "/interface bridge",
		Command: "set",
		Filter: map[string]string{
			"bridge":    bridge,
			"interface": iface,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetInterfaceBridgePortComment(bridge, iface, comment string) error {
	return r.Exec(setInterfaceBridgePort(bridge, iface, "comment", comment))
}
