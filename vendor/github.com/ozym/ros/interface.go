package ros

func interfaces() Command {
	return Command{
		Path:    "/interface",
		Command: "print",
		Detail:  true,
	}
}

func (r Ros) Interfaces() ([]map[string]string, error) {
	return r.List(interfaces())
}
