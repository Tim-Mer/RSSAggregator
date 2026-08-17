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

type state struct {
	configPtr *Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
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
	c.CurrentUserName = user
	return write(c)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Wrong number of arguments passed, expects username only")
	}
	fmt.Printf("User has been set to: %s", cmd.args[0])
	return s.configPtr.SetUser(cmd.args[0])
}

func (c *commands) run(s *state, cmd command) error {
	return c.cmdMap[cmd.name](s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}