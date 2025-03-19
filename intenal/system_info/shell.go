package system_info

import (
	"os"
	"os/exec"
	"strings"
)

func getShell() string {
	if shell := detectShellOnWindows(); shell != "" {
		return shell
	}
	if shell := detectShellOnUnix(); shell != "" {
		return shell
	}
	return "error with getting shell"

}

func detectShellOnUnix() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return ""
	}
	cmd := exec.Command(shell, "--version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return shell
	}
	return strings.TrimSpace(string(output))
}

func detectShellOnWindows() string {
	cmd := exec.Command("powershell", "-Command", "$PSVersionTable")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return "PowerShell " + strings.TrimSpace(string(output))
}
