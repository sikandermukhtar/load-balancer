// ### config/config.go
// Handles loading configuration from a JSON file.

package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ListenAddr string   `json:"listen_addr"`
	Backends   []string `json:"backends"`
	Algorithm  string   `json:"algorithm"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}