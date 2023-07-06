package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/julian776/Command-Line-Task-Manager/router"
	"github.com/julian776/Command-Line-Task-Manager/tasks/repositories"
	"github.com/julian776/Command-Line-Task-Manager/tasks/repositories/settings"
	"github.com/julian776/Command-Line-Task-Manager/tasks/services"
)

func ReadCommand(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

const (
	loadError = "can not load github.com/julian776/Command-Line-Task-Manager"
)

func Setup() *router.Router {
	settings, err := settings.LoadSettings()
	if err != nil {
		fmt.Println(loadError)
	}

	// Verify tasks file exists
	if _, err := os.Stat(settings.FilePath); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(settings.FilePath)
		if err != nil {
			fmt.Println(loadError)
		}
		file.Write([]byte("{}"))
	}

	// Setup Tasks Repo
	repo := repositories.NewTasksRepository(settings)

	// Setup Tasks Service
	service := services.NewTasksService(repo)
	// Setup Router
	router := router.NewRouter(service)

	return router
}
