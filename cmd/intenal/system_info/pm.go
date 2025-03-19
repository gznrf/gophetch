package system_info

import (
	"os"
	"os/exec"
)

func getPackageManager() string {
	name := detectPackageManagerByEnv()
	if name == "" {
		name = detectByReleaseFile()
	}
	return name
}

func detectPackageManagerByEnv() string {
	managers := map[string]string{
		"brew":   "Homebrew",
		"apt":    "APT",
		"dnf":    "DNF",
		"pacman": "Pacman",
		"zypper": "Zypper",
		"emerge": "Portage",
		"xbps":   "XBPS",
		"apk":    "Alpine Package Keeper",
	}

	for cmd, name := range managers {
		if _, err := exec.LookPath(cmd); err == nil {
			return name, true
		}
	}

	return ""
}

func detectByReleaseFile() string {
	releaseFiles := map[string]string{
		"/etc/debian_version": "APT",
		"/etc/arch-release":   "Pacman",
		"/etc/fedora-release": "DNF",
		"/etc/SuSE-release":   "Zypper",
	}

	for path, manager := range releaseFiles {
		if _, err := os.Stat(path); err == nil {
			return manager
		}
	}

	return "Error of founding pm"
}
