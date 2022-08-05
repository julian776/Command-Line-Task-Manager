package services

import (
	"errors"
	"fmt"
	"strings"
	"toDoList/tasks/models"
	"toDoList/tasks/repositories"
)

type TasksService struct {
	tasksRepository repositories.TasksRepository
}

func (service TasksService) AddTask(params []string) (string, error) {
	title := params[0]
	desc := strings.Join(params[1:], " ")
	if title == "" {
		return "ERROR: ", errors.New("Not possible to create a task with empty title")
	}
	if desc == "" {
		return "ERROR: ", errors.New("Not possible to create a task with empty description")
	}
	taskToAdd := models.Task{
		Title:       title,
		Description: desc,
	}
	service.tasksRepository.Save(taskToAdd)
	return "Task created", nil
}

func (service TasksService) UpdateDescription(params []string) (string, error) {
	title := params[0]
	desc := strings.Join(params[1:], " ")
	task, err := service.tasksRepository.FindByTitle(title)
	if err != nil {
		task.ChangeDescription(desc)
		service.tasksRepository.Save(task)
		return "Task Updated", nil
	} else {
		return "", errors.New("Can not find a task with tittle " + title)
	}
}

func (service TasksService) FindTask(params []string) (string, error) {
	fmt.Println(params[0])
	task, err := service.tasksRepository.FindByTitle(params[0])
	if err != nil {
		return "ERROR: ", err
	} else {
		return task.String(), nil
	}
}

func (service TasksService) PrintAllTasks(_ []string) (string, error) {
	tasks := service.tasksRepository.FindAll()
	for _, task := range tasks {
		fmt.Println(task)
	}
	return "", nil
}

func (service *TasksService) SetupRepository(repository repositories.TasksRepository) {
	service.tasksRepository = repository
}
