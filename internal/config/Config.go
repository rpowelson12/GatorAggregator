package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()
	return filepath.Join(path, configFileName), err
}

func Read() (Config, error) {
	homeDir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(homeDir)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(cfg Config) error {
	// 1. Marshal the Config to JSON
	json, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	// 2. Get the config file path
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	// 3. Write the JSON to the file
	return os.WriteFile(configPath, json, 0644)
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}
