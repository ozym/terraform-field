package ros

import (
	"strconv"
)

func radiusIncoming() Command {
	return Command{
		Path:    "/radius incoming",
		Command: "print",
	}
}

func (r Ros) RadiusIncoming() (map[string]string, error) {
	return r.Values(radiusIncoming())
}

func setRadiusIncomingAccept(accept bool) Command {
	return Command{
		Path:    "/radius incoming",
		Command: "set",
		Params: map[string]string{
			"accept": FormatBool(accept),
		},
	}
}

func (r Ros) SetRadiusIncomingAccept(accept bool) error {
	return r.Exec(setRadiusIncomingAccept(accept))
}

func setRadiusIncomingPort(port string) Command {
	return Command{
		Path:    "/radius incoming",
		Command: "set",
		Params: map[string]string{
			"port": port,
		},
	}
}
func (r Ros) SetRadiusIncomingPort(port int) error {
	return r.Exec(setRadiusIncomingPort(strconv.Itoa(port)))
}
