package ros

import (
	"strconv"
)

func ipService(name string) Command {
	return Command{
		Path:    "/ip service",
		Command: "print",
		Filter: map[string]string{
			"name": name,
		},
		Detail: true,
	}
}

func (r Ros) IpService(name string) (map[string]string, error) {
	return r.First(ipService(name))
}

func setIpService(name, key, value string) Command {
	return Command{
		Path:    "/ip service",
		Command: "set",
		Filter: map[string]string{
			"name": name,
		},
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetIpServiceDisabled(name string, disabled bool) error {
	return r.Exec(setIpService(name, "disabled", FormatBool(disabled)))
}

func (r Ros) SetIpServicePort(name string, port int) error {
	return r.Exec(setIpService(name, "port", strconv.Itoa(port)))
}

func (r Ros) SetIpServiceAddress(name, address string) error {
	return r.Exec(setIpService(name, "address", address))
}
