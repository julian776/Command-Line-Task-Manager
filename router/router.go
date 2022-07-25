package router

import (
	"fmt"
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
	case "show":
		task, err := router.tasksService.FindTask(operation[1])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(task)
		}
	case "add":
		_, err := router.tasksService.AddTask(operation[1], operation[2])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Added new task")
		}
	}
}

func (router *Router) SetupService(service services.TasksService) {
	router.tasksService = service
}
