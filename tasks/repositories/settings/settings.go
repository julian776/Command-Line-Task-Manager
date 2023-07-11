package settings

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	fileSettings = "settings.json"
)

type Settings struct {
	FilePath string `json:"filePath,omitempty"`
}

func NewSettings(filePath string) *Settings {
	return &Settings{
		filePath,
	}
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

func UpdateSettings(settingsToUpdate Settings) (Settings, error) {
	data, err := json.Marshal(settingsToUpdate)
	if err != nil {
		return Settings{}, err
	}
	err = os.WriteFile(fileSettings, data, os.ModePerm)
	if err != nil {
		return Settings{}, errors.New("can not create settings file, you are using root user?")
	}

	return settingsToUpdate, nil
}
