package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

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
func WriteConfig(config_path string, config any) error {
	folder := path.Dir(config_path)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Println("Creating folder", folder)
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}
	}

	jsonData, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(config_path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
