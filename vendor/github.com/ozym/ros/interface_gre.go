package ros

func interfaceGres() Command {
	return Command{
		Path:    "/interface gre",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) InterfaceGres() ([]map[string]string, error) {
	return r.List(interfaceGres())
}

func interfaceGre(address string) Command {
	return Command{
		Path:    "/interface gre",
		Command: "print",
		Filter: map[string]string{
			"remote-address": address,
		},
		Detail: true,
	}
}

func (r Ros) InterfaceGre(address string) (map[string]string, error) {
	return r.First(interfaceGre(address))
}

func setInterfaceGre(address, key, value string) Command {
	return Command{
		Path:    "/interface gre",
		Command: "set",
		Filter: map[string]string{
			"remote-address": address,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetInterfaceGreName(address, name string) error {
	return r.Exec(setInterfaceGre(address, "name", name))
}
func (r Ros) SetInterfaceGreComment(address, comment string) error {
	return r.Exec(setInterfaceGre(address, "comment", comment))
}
func (r Ros) SetInterfaceGreMtu(address, mtu string) error {
	return r.Exec(setInterfaceGre(address, "mtu", mtu))
}
func (r Ros) SetInterfaceGreKeepalive(address, alive string) error {
	return r.Exec(setInterfaceGre(address, "keepalive", alive))
}
func (r Ros) SetInterfaceGreClampTcpMss(address string, clamp bool) error {
	return r.Exec(setInterfaceGre(address, "clamp-tcp-mss", FormatBool(clamp)))
}
func (r Ros) SetInterfaceGreAllowFastPath(address string, allow bool) error {
	return r.Exec(setInterfaceGre(address, "allow-fast-path", FormatBool(allow)))
}
