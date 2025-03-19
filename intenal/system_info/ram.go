package system_info

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

func getRAM() string {
	ram := detectRAM()
	return ram
}

func detectRAM() string {
	v, _ := mem.VirtualMemory()
	return fmt.Sprintf("%.2f / %.2f GB\n", float64(v.Available)/1e9, float64(v.Total)/1e9)
}
