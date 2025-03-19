package system_info

import (
	"bufio"
	"os/exec"
	"strings"
)

func getGPU() string {
	if gpu := detectMacGPUOnLinux(); gpu != "" {
		return gpu
	}
	if gpu := detectMacGPUOnMac(); gpu != "" {
		return gpu
	}
	if gpu := detectMacGPUOnWindows(); gpu != "" {
		return gpu
	}
	return "error with getting gpu"
}

func detectMacGPUOnMac() string {
	cmd := exec.Command("system_profiler", "SPDisplaysDataType")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Chipset Model:") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}

func detectMacGPUOnWindows() string {
	cmd := exec.Command("powershell", "Get-WmiObject Win32_VideoController | Select-Object Name,AdapterRAM")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(output)
}

func detectMacGPUOnLinux() string {
	cmd := exec.Command("lspci")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(output)
}
