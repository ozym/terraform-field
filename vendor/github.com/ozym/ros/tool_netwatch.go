package ros

func toolNetwatch(host string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "print",
		Filter: map[string]string{
			"host": host,
		},
		Detail: true,
	}
}

func (r Ros) ToolNetwatch(host string) (map[string]string, error) {
	return r.First(toolNetwatch(host))
}

func addToolNetwatch(host, up_script, down_script, interval, timeout string, disabled bool) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "add",
		Params: map[string]string{
			"host":        host,
			"up-script":   up_script,
			"down-script": down_script,
			"interval":    interval,
			"timeout":     timeout,
			"disabled":    FormatBool(disabled),
		},
	}
}
func (r Ros) AddToolNetwatch(host, up_script, down_script, interval, timeout string, disabled bool) error {
	return r.Exec(addToolNetwatch(host, up_script, down_script, interval, timeout, disabled))
}

func removeToolNetwatch(host string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "remove",
		Filter: map[string]string{
			"host": host,
		},
	}
}
func (r Ros) RemoveToolNetwatch(host string) error {
	return r.Exec(removeToolNetwatch(host))
}

func setToolNetwatchUpScript(host, up_script string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "set",
		Filter: map[string]string{
			"host": host,
		},
		Params: map[string]string{
			"up-script": up_script,
		},
	}
}
func (r Ros) SetToolNetwatchUpScript(host, up_script string) error {
	return r.Exec(setToolNetwatchUpScript(host, up_script))
}

func setToolNetwatchDownScript(host, down_script string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "set",
		Filter: map[string]string{
			"host": host,
		},
		Params: map[string]string{
			"down-script": down_script,
		},
	}
}
func (r Ros) SetToolNetwatchDownScript(host, down_script string) error {
	return r.Exec(setToolNetwatchUpScript(host, down_script))
}

func setToolNetwatchInterval(host, interval string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "set",
		Filter: map[string]string{
			"host": host,
		},
		Params: map[string]string{
			"interval": interval,
		},
	}
}
func (r Ros) SetToolNetwatchInterval(host, interval string) error {
	return r.Exec(setToolNetwatchInterval(host, interval))
}

func setToolNetwatchTimeout(host, timeout string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "set",
		Filter: map[string]string{
			"host": host,
		},
		Params: map[string]string{
			"timeout": timeout,
		},
	}
}
func (r Ros) SetToolNetwatchTimeout(host, timeout string) error {
	return r.Exec(setToolNetwatchTimeout(host, timeout))
}

func setToolNetwatchDisabled(host string, disabled bool) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "set",
		Filter: map[string]string{
			"host": host,
		},
		Params: map[string]string{
			"disabled": FormatBool(disabled),
		},
	}
}
func (r Ros) SetToolNetwatchDisabled(host string, disabled bool) error {
	return r.Exec(setToolNetwatchDisabled(host, disabled))
}

func setToolNetwatchComment(host, comment string) Command {
	return Command{
		Path:    "/tool netwatch",
		Command: "set",
		Filter: map[string]string{
			"host": host,
		},
		Params: map[string]string{
			"comment": comment,
		},
	}
}
func (r Ros) SetToolNetwatchComment(host, comment string) error {
	return r.Exec(setToolNetwatchComment(host, comment))
}
