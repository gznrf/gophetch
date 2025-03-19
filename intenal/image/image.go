package image

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gznrf/gophetch.git/intenal/config"
)

type Photo struct {
	conf      *config.Config
	photoPath string
}

func NewPhoto(conf *config.Config) *Photo {
	return &Photo{
		conf: conf,
	}
}

func (photo *Photo) GetPathFromConfig() {
	photo.photoPath = photo.conf.PhotoPath
}

func (photo *Photo) Print() {
	photoPath := photo.photoPath

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Ошибка при получении домашней директории: " + err.Error())
	}

	fullPath := filepath.Join(homeDir, photoPath)
	cmd := exec.Command("viu", "-h", "10", "-w", "30", fullPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
