package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/julian776/Command-Line-Task-Manager/router"
	"github.com/julian776/Command-Line-Task-Manager/tasks/services"
	"github.com/julian776/Command-Line-Task-Manager/tasks/settings"
)

func ReadCommand(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

const (
	loadError = "can not load github.com/julian776/Command-Line-Task-Manager"
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
