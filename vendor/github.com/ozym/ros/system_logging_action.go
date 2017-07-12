package ros

import (
	"strconv"
)

func systemLoggingAction(name string) Command {
	return Command{
		Path:    "/system logging action",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) SystemLoggingAction(name string) (map[string]string, error) {
	return r.Values(systemLoggingAction(name))
}

func setSystemLoggingAction(name, key, value string) Command {
	return Command{
		Path:    "/system logging action",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetSystemLoggingActionTarget(name, target string) error {
	return r.Exec(setSystemLoggingAction(name, "target", target))
}
func (r Ros) SetSystemLoggingActionRemote(name, remote string) error {
	return r.Exec(setSystemLoggingAction(name, "remote", remote))
}
func (r Ros) SetSystemLoggingActionRemotePort(name string, port int) error {
	return r.Exec(setSystemLoggingAction(name, "remote-port", strconv.Itoa(port)))
}
func (r Ros) SetSystemLoggingActionSrcAddress(name string, address string) error {
	return r.Exec(setSystemLoggingAction(name, "src-address", address))
}
func (r Ros) SetSystemLoggingActionBsdSyslog(name string, bsd bool) error {
	return r.Exec(setSystemLoggingAction(name, "bsd-syslog", FormatBool(bsd)))
}
func (r Ros) SetSystemLoggingActionSyslogSeverity(name, severity string) error {
	return r.Exec(setSystemLoggingAction(name, "syslog-severity", severity))
}
func (r Ros) SetSystemLoggingActionSyslogFacility(name, facility string) error {
	return r.Exec(setSystemLoggingAction(name, "syslog-facility", facility))
}
