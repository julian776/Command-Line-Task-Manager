package repositories

import "toDoList/tasks/models"

type TasksRepository struct {
	tasks map[string]models.Task
}

func (repo TasksRepository) addTask(task models.Task) {

}
