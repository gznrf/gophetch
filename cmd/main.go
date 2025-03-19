package main

import (
	"github.com/gznrf/gophetch.git/intenal/config"
	"github.com/gznrf/gophetch.git/intenal/image"
	"github.com/gznrf/gophetch.git/intenal/system_info"
)

func fetch() {
	conf := config.Config{}
	conf.Load()

	image := image.NewPhoto(&conf)
	image.GetPathFromConfig()

	sysInfo := system_info.NewSystemInfo(&conf)
	sysInfo.GetByConfig()

	image.Print()
	sysInfo.Print()
}

func main() {
	fetch()
}
