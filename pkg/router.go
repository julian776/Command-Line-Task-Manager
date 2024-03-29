package pkg

import (
	"fmt"

	"github.com/julian776/Command-Line-Task-Manager/pkg/commands"
	"github.com/julian776/Command-Line-Task-Manager/pkg/tasks"
)

const (
	NOT_COMMAND_MSG = "unknown command"
)

type Router struct {
	tasksService *tasks.TasksService
}

type fn func(cmd *commands.Command) (string, error)

func (router *Router) Router(cmd *commands.Command) {
	routes := map[string]fn{
		"init": router.tasksService.Initialize,
		"help": router.tasksService.PrintFullDocs,
		"ls":   router.tasksService.PrintAllTasks,
		"show": router.tasksService.FindTask,
		"add":  router.tasksService.AddTask,
		"done": router.tasksService.CompleteTask,
	}

	if handler, exists := routes[cmd.CmdType]; exists {
		response, err := handler(cmd)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return
		}
		fmt.Println(response)
		return
	}
	fmt.Println(NOT_COMMAND_MSG + " " + cmd.CmdType)
}

func NewRouter(service *tasks.TasksService) *Router {
	return &Router{
		service,
	}
}
