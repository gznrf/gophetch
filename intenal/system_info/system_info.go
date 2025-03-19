package system_info

import (
	"fmt"

	"github.com/gznrf/gophetch.git/intenal/config"
)

type SystemInfo struct {
	conf *config.Config
	//Photo
	PhotoPath string
	//System
	OS             string
	Kernel         string
	PackageManager string
	Shell          string
	Terminal       string
	CPU            string
	GPU            string
	RAM            string
}

func NewSystemInfo(conf *config.Config) *SystemInfo {
	return &SystemInfo{
		conf: conf,
	}
}

// Сорян за эти функции самому стыдно
func (sysInf *SystemInfo) GetByConfig() {
	conf := sysInf.conf
	count := 0
	if conf.OS {
		sysInf.OS = getOS()
		count++
	}
	if conf.Kernel {
		sysInf.Kernel = getKernel()
		count++
	}
	if conf.PackageManager {
		sysInf.PackageManager = getPackageManager()
		count++
	}
	if conf.Shell {
		sysInf.Shell = getShell()
		count++
	}
	if conf.Terminal {
		sysInf.Terminal = getTerminal()
		count++
	}
	if conf.CPU {
		sysInf.CPU = getCPU()
		count++
	}
	if conf.GPU {
		sysInf.GPU = getGPU()
		count++
	}
	if conf.RAM {
		sysInf.RAM = getRAM()
		count++
	}
	if count < 1 {
		fmt.Println("check your config ~/.config/gophetch/config.json")
	}
}

func (sysInf *SystemInfo) Print() {
	if sysInf.OS != "" {
		fmt.Println(sysInf.OS)
	}
	if sysInf.Kernel != "" {
		fmt.Println(sysInf.Kernel)
	}
	if sysInf.PackageManager != "" {
		fmt.Println(sysInf.PackageManager)
	}
	if sysInf.Shell != "" {
		fmt.Println(sysInf.Shell)
	}
	if sysInf.Terminal != "" {
		fmt.Println(sysInf.Terminal)
	}
	if sysInf.CPU != "" {
		fmt.Println(sysInf.CPU)
	}
	if sysInf.GPU != "" {
		fmt.Println(sysInf.GPU)
	}
	if sysInf.RAM != "" {
		fmt.Println(sysInf.RAM)
	}
}
