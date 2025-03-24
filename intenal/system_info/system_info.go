package system_info

import (
	"fmt"

	"atomicgo.dev/cursor"
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
	conf := sysInf.conf
	cursor.Right(22)
	cursor.Up(9)
	isParamExists("OS:", sysInf.OS, conf.OSColor)
	cursor.Right(22)
	isParamExists("Kernel:", sysInf.Kernel, conf.KernelColor)
	cursor.Right(22)
	isParamExists("PM:", sysInf.PackageManager, conf.PackageManagerColor)
	cursor.Right(22)
	isParamExists("Shell", sysInf.Shell, conf.ShellColor)
	cursor.Right(22)
	isParamExists("Term:", sysInf.Terminal, conf.TerminalColor)
	cursor.Right(22)
	isParamExists("CPU", sysInf.CPU, conf.CPUColor)
	cursor.Right(22)
	isParamExists("GPU", sysInf.GPU, conf.GPUColor)
	cursor.Right(22)
	isParamExists("RAM", sysInf.RAM, conf.RAMColor)

}

func isParamExists(paramDisc, param, color string) {
	if color == "" {
		color = "#5cb7ec"
	}
	if param != "" {
		fmt.Println(paramDisc, colorText(param, color))
	}
}
