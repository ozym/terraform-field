package ros

func user(name string) Command {
	return Command{
		Path:    "/user",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) User(name string) (map[string]string, error) {
	return r.First(user(name))
}

func addUser(name, group, password string) Command {
	return Command{
		Path:    "/user",
		Command: "add",
		Params: map[string]string{
			"name":  name,
			"group": group,
		},
		Extra: map[string]string{
			"password": password,
		},
	}
}
func (r Ros) AddUser(name, group, password string) error {
	if name != "admin" {
		return r.Exec(addUser(name, group, password))
	}
	return nil
}

func removeUser(name string) Command {
	return Command{
		Path:    "/user",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
	}
}
func (r Ros) RemoveUser(name string) error {
	if name != "admin" {
		return r.Exec(removeUser(name))
	}
	return nil
}

func setUser(name, key, value string) Command {
	return Command{
		Path:    "/user",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetUserGroup(name, group string) error {
	if name != "admin" {
		return r.Exec(setUser(name, "group", group))
	}
	return nil
}
func (r Ros) SetUserComment(name, comment string) error {
	return r.Exec(setUser(name, "comment", comment))
}
