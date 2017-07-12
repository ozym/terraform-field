package ros

func interfaceWirelesses() Command {
	return Command{
		Path:    "/interface wireless",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) InterfaceWirelesses() ([]map[string]string, error) {
	return r.List(interfaceWirelesses())
}

func interfaceWireless(name string) Command {
	return Command{
		Path:    "/interface wireless",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) InterfaceWireless(name string) (map[string]string, error) {
	return r.First(interfaceWireless(name))
}

func setInterfaceWireless(name, key, value string) Command {
	return Command{
		Path:    "/interface wireless",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetInterfaceWirelessComment(name, comment string) error {
	return r.Exec(setInterfaceWireless(name, "comment", comment))
}
