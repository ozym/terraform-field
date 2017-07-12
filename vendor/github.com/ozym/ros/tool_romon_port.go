package ros

import (
	"strconv"
)

func toolRomonPort(iface string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "print",
		Filter: map[string]string{
			"interface": iface,
		},
		Detail: true,
	}
}

func (r Ros) ToolRomonPort(iface string, legacy bool) (map[string]string, error) {
	return r.First(toolRomonPort(iface, legacy))
}

func addToolRomonPort(iface string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "add",
		Params: map[string]string{
			"interface": iface,
		},
	}
}
func (r Ros) AddToolRomonPort(iface string, legacy bool) error {
	return r.Exec(addToolRomonPort(iface, legacy))
}

func removeToolRomonPort(iface string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "remove",
		Filter: map[string]string{
			"interface": iface,
		},
	}
}
func (r Ros) RemoveToolRomonPort(iface string, legacy bool) error {
	return r.Exec(removeToolRomonPort(iface, legacy))
}

func setToolRomonPort(iface, key, value string, legacy bool) Command {
	return Command{
		Path: func() string {
			if legacy {
				return "/romon port"
			}
			return "/tool romon port"
		}(),
		Command: "set",
		Filter: map[string]string{
			"interface": iface,
		},
		Params: map[string]string{
			key: value,
		},
	}
}

func (r Ros) SetToolRomonPortSecrets(iface, secrets string, legacy bool) error {
	return r.Exec(setToolRomonPort(iface, "secrets", secrets, legacy))
}
func (r Ros) SetToolRomonPortCost(iface string, cost int, legacy bool) error {
	return r.Exec(setToolRomonPort(iface, "cost", strconv.Itoa(cost), legacy))
}
func (r Ros) SetToolRomonPortDisabled(iface string, disabled bool, legacy bool) error {
	return r.Exec(setToolRomonPort(iface, "disabled", FormatBool(disabled), legacy))
}
func (r Ros) SetToolRomonPortForbid(iface string, forbid bool, legacy bool) error {
	return r.Exec(setToolRomonPort(iface, "forbid", FormatBool(forbid), legacy))
}
