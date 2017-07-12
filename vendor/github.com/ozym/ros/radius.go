package ros

func radius() Command {
	return Command{
		Path:    "/radius",
		Command: "print",
	}
}

func (r Ros) Radius() (map[string]string, error) {
	return r.Values(radius())
}

func setRadiusAddress(address string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Params: map[string]string{
			"address": address,
		},
	}
}
func (r Ros) SetRadiusAddress(address string) error {
	return r.Exec(setRadiusAddress(address))
}

func setRadiusSecret(secret string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Params: map[string]string{
			"secret": secret,
		},
	}
}
func (r Ros) SetRadiusSecret(secret string) error {
	return r.Exec(setRadiusSecret(secret))
}

func setRadiusService(service string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Params: map[string]string{
			"service": service,
		},
	}
}
func (r Ros) SetRadiusService(service string) error {
	return r.Exec(setRadiusService(service))
}

func setRadiusSrcAddress(address string) Command {
	return Command{
		Path:    "/radius",
		Command: "set",
		Params: map[string]string{
			"src-address": address,
		},
	}
}
func (r Ros) SetRadiusSrcAddress(address string) error {
	return r.Exec(setRadiusSrcAddress(address))
}
