package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/JustinTimperio/osinfo"
	"github.com/gznrf/gophetch.git/cmd/intenal/config"
	"github.com/jaypipes/ghw"
	. "github.com/klauspost/cpuid/v2"
)

func detectPackageManager() string {
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
			return name
		}
	}

	return "Не удалось определить менеджер пакетов"
}
func fetch() {
	conf := config.Config{}
	conf.Load()
	//TODO sysInfo.GetInfo(*Config)
}

func main() {

	imagePath := "./gopher_image_resized.png"

	cmd := exec.Command("viu", "-h", "10", "-w", "30", "-n", imagePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic("Ошибка при выполнении команды: " + err.Error())
	}
	er := config.Config{}
	er.Load()
	fmt.Println(er)
	fmt.Println(runtime.GOOS)
	oss := osinfo.GetVersion()
	name := oss.Name
	fmt.Println(name)
	kernel := osinfo.GetVersion().Name
	fmt.Print(kernel)
	fmt.Println(detectPackageManager())

	terminal := os.Getenv("TERM")
	fmt.Println(runtime.NumCPU())
	fmt.Print(terminal)
	fmt.Println("Name:", CPU.BrandName)
	// Получаем информацию о GPU
	gpu, err := ghw.GPU()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	for _, device := range gpu.GraphicsCards {
		fmt.Println("Видеокарта:", device.DeviceInfo.Address)
	} // СДЕЛАТЬ ВИД?ХУ И ТД
}
