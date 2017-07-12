package ros

func systemResource() Command {
	return Command{
		Path:    "/system resource",
		Command: "print",
	}
}

func (r Ros) SystemResource() (map[string]string, error) {
	return r.Values(systemResource())
}
