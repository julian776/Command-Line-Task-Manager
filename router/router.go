package router

import (
	"fmt"

	"github.com/julian776/Command-Line-Task-Manager/tasks/services"
)

type Router struct {
	tasksService services.TasksService
}

type fn func([]string) (string, error)

func (router Router) Router(operation []string) {
	routes := map[string]fn{
		"help": router.tasksService.PrintFullDocs,
		"ls":   router.tasksService.PrintAllTasks,
		"show": router.tasksService.FindTask,
		"add":  router.tasksService.AddTask,
		"done": router.tasksService.CompleteTask,
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
