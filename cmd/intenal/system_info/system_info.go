package system_info

import (
	"github.com/gznrf/gophetch.git/cmd/intenal/config"
)

type SystemInfo struct {
	conf config.Config
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

func (sysInf *SystemInfo) getByConfig() {
	conf := sysInf.conf
	switch {
	case conf.OS:
		sysInf.OS = getOS()
		fallthrough
	case conf.Kernel:
		sysInf.Kernel = getKernel()
		fallthrough
	case conf.PackageManager:
		sysInf.PackageManager = getPackageManager()
		fallthrough
	case conf.Shell:
		sysInf.Shell = getShell()
		fallthrough
	case conf.Terminal:
		sysInf.Terminal = getTerminal()
		fallthrough
	case conf.CPU:
		//sysInf.CPU = getCPU()
		fallthrough
	case conf.GPU:
		//sysInf.GPU = getGPU()
		fallthrough
	case conf.RAM:
		//sysInf.RAM = getRAM()
		break
	}
}
