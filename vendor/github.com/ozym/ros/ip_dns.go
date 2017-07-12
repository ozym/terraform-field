package ros

func ipDns() Command {
	return Command{
		Path:    "/ip dns",
		Command: "print",
	}
}

func (r Ros) IpDns() (map[string]string, error) {
	return r.Values(ipDns())
}

func setIpDns(key, value string) Command {
	return Command{
		Path:    "/ip dns",
		Command: "set",
		Params: map[string]string{
			key: value,
		},
	}
}
func (r Ros) SetIpDnsServers(servers string) error {
	return r.Exec(setIpDns("servers", servers))
}
func (r Ros) SetIpDnsAllowRemoteRequests(allow bool) error {
	return r.Exec(setIpDns("allow-remote-requests", FormatBool(allow)))
}
