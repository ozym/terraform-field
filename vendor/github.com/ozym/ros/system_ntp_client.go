package ros

func systemNtpClient() Command {
	return Command{
		Path:    "/system ntp client",
		Command: "print",
	}
}

func (r Ros) SystemNtpClient() (map[string]string, error) {
	return r.Values(systemNtpClient())
}

func setSystemNtpClientEnabled(enabled bool) Command {
	return Command{
		Path:    "/system ntp client",
		Command: "set",
		Params: map[string]string{
			"enabled": FormatBool(enabled),
		},
	}
}
func (r Ros) SetSystemNtpClientEnabled(enabled bool) error {
	return r.Exec(setSystemNtpClientEnabled(enabled))
}
func setSystemNtpClientPrimaryNtp(host string) Command {
	return Command{
		Path:    "/system ntp client",
		Command: "set",
		Params: map[string]string{
			"primary-ntp": host,
		},
	}
}
func (r Ros) SetSystemNtpClientPrimaryNtp(zone string) error {
	return r.Exec(setSystemNtpClientPrimaryNtp(zone))
}
func setSystemNtpClientSecondaryNtp(host string) Command {
	return Command{
		Path:    "/system ntp client",
		Command: "set",
		Params: map[string]string{
			"secondary-ntp": host,
		},
	}
}
func (r Ros) SetSystemNtpClientSecondaryNtp(zone string) error {
	return r.Exec(setSystemNtpClientSecondaryNtp(zone))
}
