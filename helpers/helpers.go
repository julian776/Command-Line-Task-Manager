package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"toDoList/router"
	"toDoList/tasks/services"
	"toDoList/tasks/settings"
)

func ReadCommand(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

const (
	loadError = "can not load toDoList"
)

func Setup() router.Router {
	settings, err := settings.LoadSettings()
	if err != nil {
		fmt.Println(loadError)
	}

	// Verify tasks file exists
	if _, err := os.Stat(settings.FileName); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(settings.FileName)
		if err != nil {
			fmt.Println(loadError)
		}
		file.Write([]byte("{}"))
	}

	// Setup Tasks Service
	service := services.TasksService{}
	service.NewTasksService(settings)
	// Setup Router
	router := router.Router{}
	router.SetupService(service)

	return router
}
