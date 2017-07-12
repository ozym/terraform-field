package ros

func interfaceBridges() Command {
	return Command{
		Path:    "/interface bridge",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) InterfaceBridges() ([]map[string]string, error) {
	return r.List(interfaceBridges())
}

func interfaceBridge(name string) Command {
	return Command{
		Path:    "/interface bridge",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) InterfaceBridge(name string) (map[string]string, error) {
	return r.First(interfaceBridge(name))
}

func setInterfaceBridge(name, key, value string) Command {
	return Command{
		Path:    "/interface bridge",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetInterfaceBridgeComment(name, comment string) error {
	return r.Exec(setInterfaceBridge(name, "comment", comment))
}

func (r Ros) SetInterfaceBridgeProtocolMode(name, mode string) error {
	return r.Exec(setInterfaceBridge(name, "protocol-mode", mode))
}

func (r Ros) SetInterfaceBridgePriority(name, priority string) error {
	return r.Exec(setInterfaceBridge(name, "priority", priority))
}
