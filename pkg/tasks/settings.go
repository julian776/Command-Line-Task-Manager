package tasks

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const (
	fileSettings   = "settings.json"
	fileTasks      = "tasks.json"
	dirTaskManager = "Command-Line-Task-Manager"
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
	settingsPath, err := buildSettingsPath()
	if err != nil {
		return Settings{}, err
	}

	content, err := os.ReadFile(settingsPath)
	if err != nil {
		return Settings{}, err
	}

	err = json.Unmarshal(content, &settings)
	if err != nil {
		return Settings{}, err
	}

	if settings.FilePath == "" {
		return Settings{}, errors.New("not settings file, please run init command")
	}

	return settings, nil
}

func buildSettingsPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return homePath, err
	}
	path := filepath.Join(homePath, dirTaskManager, fileSettings)
	return path, nil
}

func buildTasksPath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return homePath, err
	}
	path := filepath.Join(homePath, dirTaskManager, fileTasks)
	return path, nil
}

func CreateDefaultSettingsFile() error {
	settingsPath, err := buildSettingsPath()
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(settingsPath), os.ModePerm)
	if err != nil {
		return errors.New("can not create default settings, you are using root user?")
	}

	tasksPath, err := buildTasksPath()
	if err != nil {
		return err
	}

	settings := NewSettings(tasksPath)
	fileData, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	err = os.WriteFile(settingsPath, fileData, os.ModePerm)
	if err != nil {
		return errors.New("can not create default settings, you are using root user?")
	}

	return nil
}
