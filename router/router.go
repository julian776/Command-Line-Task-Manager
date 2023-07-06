package router

import (
	"fmt"

	"github.com/julian776/Command-Line-Task-Manager/commands/models"
	"github.com/julian776/Command-Line-Task-Manager/tasks/services"
)

const (
	NOT_COMMAND_MSG = "unknown command"
)

type Router struct {
	tasksService *services.TasksService
}

type fn func(cmd *models.Command) (string, error)

func (router *Router) Router(cmd *models.Command) {
	routes := map[string]fn{
		"help": router.tasksService.PrintFullDocs,
		"ls":   router.tasksService.PrintAllTasks,
		"show": router.tasksService.FindTask,
		"add":  router.tasksService.AddTask,
		"done": router.tasksService.CompleteTask,
	}

	if handler, exists := routes[cmd.CmdType]; exists {
		response, err := handler(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(response)
		return
	}
	fmt.Println(NOT_COMMAND_MSG + " " + cmd.CmdType)
}

func NewRouter(service *services.TasksService) *Router {
	return &Router{
		service,
	}
}
