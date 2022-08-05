package router

import (
	"fmt"
	"strings"
	"toDoList/tasks/services"
)

type Router struct {
	tasksService services.TasksService
}

type fn func([]string) (string, error)

func (router Router) Router(command string) {
	operation := strings.Split(command, " ")
	routes := map[string]fn{
		"ls":   router.tasksService.PrintAllTasks,
		"show": router.tasksService.FindTask,
		"add":  router.tasksService.AddTask,
	}

	if value, exists := routes[operation[0]]; exists {
		response, err := value(operation[1:])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(response)
		}
	}
}

func (router *Router) SetupService(service services.TasksService) {
	router.tasksService = service
}
