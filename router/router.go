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
		description := strings.Join(operation[2:], " ")
		_, err := router.tasksService.AddTask(operation[1], description)
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
