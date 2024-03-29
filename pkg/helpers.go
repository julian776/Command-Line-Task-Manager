package pkg

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/julian776/Command-Line-Task-Manager/pkg/tasks"
)

func ReadCommand(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

const (
	loadError = "can not load github.com/julian776/Command-Line-Task-Manager"
)

func Setup() *Router {
	settings, err := tasks.LoadSettings()
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
	repo := tasks.NewTasksRepository(settings)

	// Setup Tasks Service
	service := tasks.NewTasksService(repo)
	// Setup Router
	router := NewRouter(service)

	return router
}
