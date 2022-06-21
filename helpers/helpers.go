package helpers

import (
	"bufio"
	"toDoList/router"
	"toDoList/tasks/repositories"
	"toDoList/tasks/services"
)

func ReadCommand(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func Setup() router.Router {
	repository := repositories.TasksRepository{}
	// Setup Tasks Service
	service := services.TasksService{}
	service.SetupRepository(repository)
	// Setup Router
	router := router.Router{}
	router.SetupService(service)

	return router
}
