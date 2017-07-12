package ros

func systemLogging(action, topics string) Command {
	return Command{
		Path:    "/system logging",
		Command: "print",
		Filter: map[string]string{
			"action": action,
			"topics": topics,
		},
		Detail: true,
	}
}

func (r Ros) SystemLogging(action, topics string) (map[string]string, error) {
	return r.Values(systemLogging(action, topics))
}

func addSystemLogging(action, topics string) Command {
	return Command{
		Path:    "/system logging",
		Command: "add",
		Params: map[string]string{
			"action": action,
			"topics": topics,
		},
	}
}
func (r Ros) AddSystemLogging(action, topics string) error {
	return r.Exec(addSystemLogging(action, topics))
}

func removeSystemLogging(action, topics string) Command {
	return Command{
		Path:    "/system logging",
		Command: "remove",
		Filter: map[string]string{
			"action": action,
			"topics": topics,
		},
	}
}
func (r Ros) RemoveSystemLogging(action, topics string) error {
	return r.Exec(removeSystemLogging(action, topics))
}

func setSystemLoggingPrefix(action, topics, prefix string) Command {
	return Command{
		Path:    "/system logging",
		Command: "set",
		Filter: map[string]string{
			"action": action,
			"topics": topics,
		},
		Params: map[string]string{
			"prefix": prefix,
		},
	}
}
func (r Ros) SetSystemLoggingPrefix(action, topics, prefix string) error {
	return r.Exec(setSystemLoggingPrefix(action, topics, prefix))
}
