package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	//Photo
	PhotoPath string `json:"photo_path"`
	//System
	OS             bool `json:"os"`
	Kernel         bool `json:"kernel"`
	PackageManager bool `json:"package_manager"`
	Shell          bool `json:"shell"`
	Terminal       bool `json:"terminal"`
	CPU            bool `json:"cpu"`
	GPU            bool `json:"gpu"`
	RAM            bool `json:"ram"`
}

func NewConfig(photoPath string,
	os,
	kernel,
	packageManager,
	shell,
	wm,
	terminal,
	cpu,
	gpu,
	ram bool) *Config {
	return &Config{
		PhotoPath:      photoPath,
		OS:             os,
		Kernel:         kernel,
		PackageManager: packageManager,
		Shell:          shell,
		Terminal:       terminal,
		CPU:            cpu,
		GPU:            gpu,
		RAM:            ram,
	}
}

func (conf *Config) Load() {
	// Получаем домашнюю директорию
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Ошибка при получении домашней директории: " + err.Error())
	}

	// Корректный путь к файлу
	configPath := filepath.Join(homeDir, ".config", "gophetch", "config.json")
	jsonFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
		panic("Error with opening file")
	}
	fmt.Println("Successfully Opened json file")
	defer jsonFile.Close()

	byteEmpValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteEmpValue, &conf)
}
