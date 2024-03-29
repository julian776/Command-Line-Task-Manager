package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type TasksRepository struct {
	settings Settings
}

func (r *TasksRepository) FindByTitle(title string) (Task, error) {
	tasks, err := getTasksFromFile(r.settings.FilePath)

	if err != nil {
		return Task{}, err
	}
	if task, ok := tasks[title]; ok {
		return task, nil
	}

	return Task{}, errors.New("can not find Task")
}

func (r *TasksRepository) Save(task Task) error {
	tasks, err := getTasksFromFile(r.settings.FilePath)
	if err != nil {
		return err
	}
	tasks[task.Title] = task
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(r.settings.FilePath, data, fs.FileMode(os.ModePerm))
	if err != nil {
		return err
	}
	return nil
}

func (r *TasksRepository) FindAll() (map[string]Task, error) {
	tasks, err := getTasksFromFile(r.settings.FilePath)
	if err != nil {
		return map[string]Task{}, err
	}
	return tasks, nil
}

func getTasksFromFile(filePath string) (tasks map[string]Task, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return map[string]Task{}, fmt.Errorf("error reading tasks: %s", err.Error())
	}
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return map[string]Task{}, fmt.Errorf("error reading tasks: %s", err.Error())
	}
	return tasks, nil
}

func NewTasksRepository(settings Settings) *TasksRepository {
	return &TasksRepository{
		settings,
	}
}
