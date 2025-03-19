package system_info

import (
	"os/exec"
	"strings"
)

func getCPU() string {
	if cpu := detectCPUOnMac(); cpu != "" {
		return cpu
	}
	if cpu := detectCPUOnLinux(); cpu != "" {
		return cpu
	}
	if cpu := detectCPUOnWindows(); cpu != "" {
		return cpu
	}
	return ""
}

func detectCPUOnMac() string {
	cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(output))
}

func detectCPUOnLinux() string {
	cmd := exec.Command("cat", "/proc/cpuinfo")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "model name") {
			return strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}
	return ""
}

func detectCPUOnWindows() string {
	cmd := exec.Command("powershell", "Get-CimInstance Win32_Processor | Select-Object Name")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) != "" && !strings.Contains(line, "Name") {
			return strings.TrimSpace(line)
		}
	}

	return ""
}
