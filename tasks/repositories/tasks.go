package repositories

import (
	"errors"
	"toDoList/tasks/models"
)

type TasksRepository struct {
	tasks map[string]models.Task
}

func (repo TasksRepository) FindByTitle(title string) (models.Task, error) {
	if task, ok := repo.tasks[title]; ok {
		return task, nil
	} else {
		return models.Task{}, errors.New("Can not find Task")
	}
}

func (repo *TasksRepository) Save(task models.Task) {
	repo.tasks[task.Title] = task
}

func (repo TasksRepository) FindAll() map[string]models.Task {
	return repo.tasks
}
