package ros

import (
	"strings"
)

func systemScript(name string) Command {
	return Command{
		Path:    "/system script",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) SystemScript(name string) (map[string]string, error) {
	return r.First(systemScript(name))
}

func addSystemScript(name, policy, source string) Command {
	return Command{
		Path:    "/system script",
		Command: "add",
		Params: map[string]string{
			"name":   name,
			"policy": policy,
			"source": source,
		},
	}
}
func (r Ros) AddSystemScript(name, policy, source string) error {
	return r.Exec(addSystemScript(name, policy, source))
}

func removeSystemScript(name string) Command {
	return Command{
		Path:    "/system script",
		Command: "remove",
		Filter: map[string]string{
			"name": name,
		},
	}
}
func (r Ros) RemoveSystemScript(name string) error {
	return r.Exec(removeSystemScript(name))
}

func setSystemScriptPolicy(name, policy string) Command {
	return Command{
		Path:    "/system script",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			"policy": policy,
		},
	}
}
func (r Ros) SetSystemScriptPolicy(name, policy string) error {
	return r.Exec(setSystemScriptPolicy(name, policy))
}

func setSystemScriptSource(name, source string) Command {
	return Command{
		Path:    "/system script",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			"source": strings.Replace(source, "\"", "\\\"", -1),
		},
	}
}
func (r Ros) SetSystemScriptSource(name, source string) error {
	return r.Exec(setSystemScriptSource(name, source))
}
