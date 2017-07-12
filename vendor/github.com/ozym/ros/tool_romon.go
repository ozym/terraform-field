package ros

func toolRomon(legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon"
			}
			return "/tool romon"
		}(),
		Command: "print",
	}
}

func (r Ros) ToolRomon(legacy bool) (map[string]string, error) {
	return r.Values(toolRomon(legacy))
}

func setToolRomon(key, value string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon"
			}
			return "/tool romon"
		}(),
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetToolRomonId(id string, legacy bool) error {
	return r.Exec(setToolRomon("id", id, legacy))
}
func (r Ros) SetToolRomonEnabled(enabled bool, legacy bool) error {
	return r.Exec(setToolRomon("enabled", FormatBool(enabled), legacy))
}
func (r Ros) SetToolRomonSecrets(secrets string, legacy bool) error {
	return r.Exec(setToolRomon("secrets", secrets, legacy))
}
