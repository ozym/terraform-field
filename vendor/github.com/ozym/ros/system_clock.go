package ros

func systemClock() Command {
	return Command{
		Path:    "/system clock",
		Command: "print",
	}
}

func (r Ros) SystemClock() (map[string]string, error) {
	return r.Values(systemClock())
}

func setSystemClockTimeZoneName(zone string) Command {
	return Command{
		Path:    "/system clock",
		Command: "set",
		Params: map[string]string{
			"time-zone-name": zone,
		},
	}
}
func (r Ros) SetSystemClockTimeZoneName(zone string) error {
	return r.Exec(setSystemClockTimeZoneName(zone))
}
func setSystemClockTimeZoneAutodetect(auto bool) Command {
	return Command{
		Path:    "/system clock",
		Command: "set",
		Params: map[string]string{
			"time-zone-autodetect": FormatBool(auto),
		},
	}
}
func (r Ros) SetSystemClockTimeZoneAutodetect(auto bool) error {
	return r.Exec(setSystemClockTimeZoneAutodetect(auto))
}
