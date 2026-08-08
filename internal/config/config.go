package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Account struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Host     string `json:"host"`
}

type Config struct {
	CurrentAccount string             `json:"current_account"`
	Language       string             `json:"language"`
	Accounts       map[string]Account `json:"accounts"`
}

func GetConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "ghx"), nil
}

func GetConfigPath() (string, error) {
	dir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "accounts.json"), nil
}

func DefaultConfig() *Config {
	return &Config{
		CurrentAccount: "",
		Language:       "en",
		Accounts:       make(map[string]Account),
	}
}

func LoadConfig() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return DefaultConfig(), err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return DefaultConfig(), nil
		}
		return DefaultConfig(), err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return DefaultConfig(), err
	}

	if cfg.Accounts == nil {
		cfg.Accounts = make(map[string]Account)
	}
	if cfg.Language == "" {
		cfg.Language = "en"
	}

	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	dir, err := GetConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}

	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	if cfg.Accounts == nil {
		cfg.Accounts = make(map[string]Account)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}
