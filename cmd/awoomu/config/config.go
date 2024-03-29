package config

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
)

type AwooConfig struct {
	CPU AwooConfigCPU `json:"cpu"`
}

type AwooConfigCPU struct {
	Speed uint32 `json:"speed"`
}

func NewAwooConfig() AwooConfig {
	return AwooConfig{
		CPU: AwooConfigCPU{
			Speed: 1000,
		},
	}
}

func ReadConfig(config AwooConfig) (AwooConfig, error) {
	u, _ := user.Current()
	path := filepath.Join(u.HomeDir, "Documents", "awoo", "config", "emu.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		if raw, err = json.Marshal(config); err != nil {
			return AwooConfig{}, err
		}
		if err = os.WriteFile(path, raw, 0755); err != nil {
			return AwooConfig{}, err
		}
	}
	err = json.Unmarshal(raw, &config)
	if err != nil {
		return AwooConfig{}, err
	}

	return config, nil
}
