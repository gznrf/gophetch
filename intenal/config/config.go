package config

import (
	"encoding/json"
	"fmt"
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

func (conf *Config) Load() {
	// Получаем домашнюю директорию
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Ошибка при получении домашней директории: " + err.Error())
	}

	// Корректный путь к файлу
	configPath := filepath.Join(homeDir, ".config", "gophetch", "config.json")
	jsonFile, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		panic("Error with opening file")
	}
	err = json.Unmarshal(jsonFile, &conf)
	if err != nil {
		fmt.Println(err)
	}
}
