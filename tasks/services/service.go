package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/julian776/Command-Line-Task-Manager/tasks/models"
	"github.com/julian776/Command-Line-Task-Manager/tasks/repositories"
	"github.com/julian776/Command-Line-Task-Manager/tasks/settings"
)

type TasksService struct {
	settings settings.Settings
}

const (
	errorStr = "ERROR"
)

func (s *TasksService) AddTask(params []string) (string, error) {
	title := params[0]
	desc := strings.Join(params[1:], " ")
	if title == "" {
		return errorStr, errors.New("not possible to create a task with empty title")
	}
	if desc == "" {
		return errorStr, errors.New("not possible to create a task with empty description")
	}
	taskToAdd := models.Task{
		Title:       title,
		Description: desc,
		IsCompleted: false,
	}
	repositories.Save(s.settings.FileName, taskToAdd)
	return "Task created", nil
}

func (s *TasksService) UpdateDescription(params []string) (string, error) {
	title := params[0]
	desc := strings.Join(params[1:], " ")
	task, err := repositories.FindByTitle(s.settings.FileName, title)
	if err != nil {
		task.ChangeDescription(desc)
		repositories.Save(s.settings.FileName, task)
		return "Task Updated", nil
	} else {
		return errorStr, errors.New("Can not find a task with tittle " + title)
	}
}

func (s *TasksService) FindTask(params []string) (string, error) {
	fmt.Println(params[0])
	task, err := repositories.FindByTitle(s.settings.FileName, params[0])
	if err != nil {
		return "", err
	} else {
		return task.String(), nil
	}
}

func (s *TasksService) PrintAllTasks(_ []string) (string, error) {
	tasks, err := repositories.FindAll(s.settings.FileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, task := range tasks {
		fmt.Println(task)
	}
	return "", nil
}

func (s *TasksService) NewTasksService(settings settings.Settings) {
	s.settings = settings
}
