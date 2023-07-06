package settings

import (
	"encoding/json"
	"os"
)

const (
	fileSettings = "settings.json"
)

type Settings struct {
	FilePath string `json:"filePath,omitempty"`
}

func LoadSettings() (settings Settings, err error) {
	content, err := os.ReadFile(fileSettings)
	if err != nil {
		return Settings{}, err
	}
	err = json.Unmarshal(content, &settings)
	if err != nil {
		return Settings{}, err
	}
	return settings, nil
}
