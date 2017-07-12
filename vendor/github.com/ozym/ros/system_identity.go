package ros

func systemIdentity() Command {
	return Command{
		Path:    "/system identity",
		Command: "print",
	}
}

func (r Ros) SystemIdentity() (map[string]string, error) {
	return r.Values(systemIdentity())
}

func setSystemIdentityName(name string) Command {
	return Command{
		Path:    "/system identity",
		Command: "set",
		Params: map[string]string{
			"name": name,
		},
	}
}

func (r Ros) SetSystemIdentityName(name string) error {
	return r.Exec(setSystemIdentityName(name))
}
