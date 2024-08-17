package config

import (
	"encoding/json"
	"os"
)

var (
	BasePath = "/var/opt/homelab-scm"
)

func init() {
	if os.Getenv("HOMELAB_SCM_BASE_PATH") != "" {
		BasePath = os.Getenv("HOMELAB_SCM_BASE_PATH")
	}
}

// Reads a JSON config file into the provided struct.
// If the file does not exist, it will be created with the provided config.
func ReadConfig(path string, config any) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return WriteConfig(path, config)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		return err
	}

	return nil
}

// Writes the provided struct to a JSON file.
func WriteConfig(path string, config any) error {
	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
