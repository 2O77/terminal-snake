package config

import (
	"encoding/json"
	"os"
	"path"
)

const (
	configDirName  = ".config/terminal-snake"
	configFileName = "config.json"
)

type Config struct {
	BestScore int
}

// Load reads the config from disk. If the file does not exist (or is broken)
// it creates a default one and returns that.
func Load() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	filePath := path.Join(homeDir, configDirName, configFileName)
	byteConfig, err := os.ReadFile(filePath)
	if err == nil {
		var cfg Config
		if json.Unmarshal(byteConfig, &cfg) == nil {
			return cfg, nil
		}
		// broken config falls through to a fresh default
	}

	cfg := Config{}
	if err := cfg.Save(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

// Save writes the config to disk, creating the directory if needed.
func (c Config) Save() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dirPath := path.Join(homeDir, configDirName)
	if err := os.MkdirAll(dirPath, 0o700); err != nil {
		return err
	}

	byteConfig, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(dirPath, configFileName), byteConfig, 0o600)
}
