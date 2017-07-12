package ros

func userAaa() Command {
	return Command{
		Path:    "/user aaa",
		Command: "print",
	}
}

func (r Ros) UserAaa() (map[string]string, error) {
	return r.Values(userAaa())
}

func setUserAaaUseRadius(enabled bool) Command {
	return Command{
		Path:    "/user aaa",
		Command: "set",
		Params: map[string]string{
			"use-radius": FormatBool(enabled),
		},
	}
}
func (r Ros) SetUserAaaUseRadius(enabled bool) error {
	return r.Exec(setUserAaaUseRadius(enabled))
}

func setUserAaaAccounting(enabled bool) Command {
	return Command{
		Path:    "/user aaa",
		Command: "set",
		Params: map[string]string{
			"accounting": FormatBool(enabled),
		},
	}
}
func (r Ros) SetUserAaaAccounting(enabled bool) error {
	return r.Exec(setUserAaaAccounting(enabled))
}

func setUserAaaInterimUpdate(update string) Command {
	return Command{
		Path:    "/user aaa",
		Command: "set",
		Params: map[string]string{
			"interim-update": update,
		},
	}
}
func (r Ros) SetUserAaaInterimUpdate(update string) error {
	return r.Exec(setUserAaaInterimUpdate(update))
}

func setUserAaaDefaultGroup(group string) Command {
	return Command{
		Path:    "/user aaa",
		Command: "set",
		Params: map[string]string{
			"default-group": group,
		},
	}
}
func (r Ros) SetUserAaaDefaultGroup(group string) error {
	return r.Exec(setUserAaaDefaultGroup(group))
}

func setUserAaaExcludeGroups(groups string) Command {
	return Command{
		Path:    "/user aaa",
		Command: "set",
		Params: map[string]string{
			"exclude-groups": groups,
		},
	}
}
func (r Ros) SetUserAaaExcludeGroups(groups string) error {
	return r.Exec(setUserAaaExcludeGroups(groups))
}
