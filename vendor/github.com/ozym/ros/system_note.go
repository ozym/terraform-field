package ros

func systemNote() Command {
	return Command{
		Path:    "/system note",
		Command: "print",
	}
}

func (r Ros) SystemNote() (map[string]string, error) {
	return r.Values(systemNote())
}

func setSystemNote(note string) Command {
	return Command{
		Path:    "/system note",
		Command: "set",
		Params: map[string]string{
			"note": note,
		},
	}
}

func (r Ros) SetSystemNote(note string) error {
	return r.Exec(setSystemNote(note))
}

func setSystemNoteShowAtLogin(show bool) Command {
	return Command{
		Path:    "/system note",
		Command: "set",
		Params: map[string]string{
			"show-at-login": FormatBool(show),
		},
	}
}
func (r Ros) SetSystemNoteShowAtLogin(show bool) error {
	return r.Exec(setSystemNoteShowAtLogin(show))
}
