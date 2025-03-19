package system_info

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getShell() string {
	shell := detectShell()
	return shell

}

func detectShell() string {
	cmd := exec.Command("ps", "-p", fmt.Sprintf("%d", os.Getppid()), "-o", "comm=")
	output, err := cmd.Output()
	if err != nil {
		return "error with getting shell"
	}
	return strings.TrimSpace(string(output))
}
