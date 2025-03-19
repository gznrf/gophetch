package system_info

import "os"

func getTerminal() string {
	term := detectTerminal()
	return term
}

func detectTerminal() string {
	terminal := os.Getenv("TERM")
	if terminal == "" {
		return "error with getting term"
	}
	return terminal
}
