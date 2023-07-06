package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/julian776/Command-Line-Task-Manager/tasks/models"
)

type TasksRepository struct {
	FilePath string
}

func (r *TasksRepository) FindByTitle(title string) (models.Task, error) {
	tasks, err := getTasksFromFile(r.FilePath)
	if err != nil {
		return models.Task{}, err
	}
	if task, ok := tasks[title]; ok {
		return task, nil
	} else {
		return models.Task{}, errors.New("can not find Task")
	}
}

func (r *TasksRepository) Save(task models.Task) error {
	tasks, err := getTasksFromFile(r.FilePath)
	if err != nil {
		return err
	}
	tasks[task.Title] = task
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(r.FilePath, data, fs.FileMode(os.ModePerm))
	if err != nil {
		return err
	}
	return nil
}

func (r *TasksRepository) FindAll() (map[string]models.Task, error) {
	tasks, err := getTasksFromFile(r.FilePath)
	if err != nil {
		return map[string]models.Task{}, err
	}
	return tasks, nil
}

func getTasksFromFile(filePath string) (tasks map[string]models.Task, err error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return map[string]models.Task{}, fmt.Errorf("error reading tasks: %s", err.Error())
	}
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return map[string]models.Task{}, fmt.Errorf("error reading tasks: %s", err.Error())
	}
	return tasks, nil
}

func NewTasksRepository(filePath string) *TasksRepository {
	return &TasksRepository{
		filePath,
	}
}
