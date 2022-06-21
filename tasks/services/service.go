package services

import (
	"errors"
	"fmt"
	"toDoList/tasks/models"
	"toDoList/tasks/repositories"
)

type TasksService struct {
	tasksRepository repositories.TasksRepository
}

func (service TasksService) UpdateDescription(title string, desc string) (string, error) {
	task, err := service.tasksRepository.FindByTitle(title)
	if err != nil {
		task.ChangeDescription(desc)
		service.tasksRepository.Save(task)
		return "Task Updated", nil
	} else {
		return "", errors.New("Can not find a task with tittle " + title)
	}
}

func (service TasksService) FindTask(title string) (models.Task, error) {
	task, err := service.tasksRepository.FindByTitle(title)
	if err != nil {
		return task, nil
	} else {
		return models.Task{}, err
	}
}

func (service TasksService) PrintAllTasks() {
	tasks := service.tasksRepository.FindAll()
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func (service *TasksService) SetupRepository(repository repositories.TasksRepository) {
	service.tasksRepository = repository
}
