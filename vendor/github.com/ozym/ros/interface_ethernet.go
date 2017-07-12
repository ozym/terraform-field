package ros

func interfaceEthernets() Command {
	return Command{
		Path:    "/interface ethernet",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) InterfaceEthernets() ([]map[string]string, error) {
	return r.List(interfaceEthernets())
}

func interfaceEthernet(name string) Command {
	return Command{
		Path:    "/interface ethernet",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) InterfaceEthernet(name string) (map[string]string, error) {
	return r.First(interfaceEthernet(name))
}

func setInterfaceEthernet(name, key, value string) Command {
	return Command{
		Path:    "/interface ethernet",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetInterfaceEthernetComment(name, comment string) error {
	return r.Exec(setInterfaceEthernet(name, "comment", comment))
}
func (r Ros) SetInterfaceEthernetMtu(name, mtu string) error {
	return r.Exec(setInterfaceEthernet(name, "mtu", mtu))
}
