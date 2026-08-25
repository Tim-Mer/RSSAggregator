package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

func Read() (Config, error) {
	config := Config{}

	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	dat, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(dat, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func write(c Config) error {
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	dat, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFile, dat, 0644)
	return err
}

func (c Config) SetUser(user string) error {
	fmt.Printf("User has been set to: %s\n", user)
	c.CurrentUserName = user
	return write(c)
}
