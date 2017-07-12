package ros

func snmp() Command {
	return Command{
		Path:    "/snmp",
		Command: "print",
	}
}

func (r Ros) Snmp() (map[string]string, error) {
	return r.Values(snmp())
}

func setSnmp(key, value string) Command {
	return Command{
		Path:    "/snmp",
		Command: "set",
		Params: map[string]string{
			"key": value,
		},
	}
}

func (r Ros) SetSnmpEnabled(enabled bool) error {
	return r.Exec(setSnmp("enabled", FormatBool(enabled)))
}
func (r Ros) SetSnmpLocation(location string) error {
	return r.Exec(setSnmp("location", location))
}
func (r Ros) SetSnmpContact(contact string) error {
	return r.Exec(setSnmp("contact", contact))
}
