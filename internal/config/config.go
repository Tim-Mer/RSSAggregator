package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL string `json:"db_url"`
}

func Read() (Config, error) {
	config := Config{}
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	dat, err := os.ReadFile(filepath.Join(home, ".gatorconfig.json"))
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(dat, &config)
	if err != nil {
		return Config{}, err
	}
	fmt.Println(config.DbURL)
	return config, nil
}
