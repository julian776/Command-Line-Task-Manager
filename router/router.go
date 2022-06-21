package router

import (
	"strings"
	"toDoList/tasks/services"
)

type Router struct {
	tasksService services.TasksService
}

func (router Router) Router(command string) {
	operation := strings.Split(command, " ")

	switch operation[0] {
	case "ls":
		router.tasksService.PrintAllTasks()
	}
}

func (router *Router) SetupService(service services.TasksService) {
	router.tasksService = service
}
