package system_info

import "runtime"

func getKernel() string {
	kernel := runtime.GOOS
	return kernel

}
