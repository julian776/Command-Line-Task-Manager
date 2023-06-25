package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/julian776/Command-Line-Task-Manager/tasks/models"
)

func FindByTitle(filePath string, title string) (models.Task, error) {
	tasks, err := getTasksFromFile(filePath)
	if err != nil {
		return models.Task{}, err
	}
	if task, ok := tasks[title]; ok {
		return task, nil
	} else {
		return models.Task{}, errors.New("can not find Task")
	}
}

func Save(filePath string, task models.Task) error {
	tasks, err := getTasksFromFile(filePath)
	if err != nil {
		return err
	}
	tasks[task.Title] = task
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, data, fs.FileMode(os.ModePerm))
	if err != nil {
		return err
	}
	return nil
}

func FindAll(filePath string) (map[string]models.Task, error) {
	tasks, err := getTasksFromFile(filePath)
	if err != nil {
		return map[string]models.Task{}, err
	}
	return tasks, nil
}

func getTasksFromFile(path string) (tasks map[string]models.Task, err error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return map[string]models.Task{}, fmt.Errorf("error reading tasks: %s", err.Error())
	}
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return map[string]models.Task{}, fmt.Errorf("error reading tasks: %s", err.Error())
	}
	return tasks, nil
}
