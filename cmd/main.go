package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gznrf/gophetch.git/cmd/intenal/config"
)

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
}
