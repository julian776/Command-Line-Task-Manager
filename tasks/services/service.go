package services

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	cmdsModel "github.com/julian776/Command-Line-Task-Manager/commands/models"
	"github.com/julian776/Command-Line-Task-Manager/tasks/models"
	"github.com/julian776/Command-Line-Task-Manager/tasks/repositories"
	"github.com/julian776/Command-Line-Task-Manager/tasks/repositories/settings"
)

type TasksService struct {
	TasksRepo *repositories.TasksRepository
}

const (
	errorStr = "ERROR"

	initializeErrorStr = "not posible to initialize Command-Line-Task-Manager, you are using root user?"

	FULL_DOCS = `
ls - List all your tasks.
	Example:
		Command-Line-Task-Manager ls

add - Add a Task
	The add command allows you to add
	a new task to your task list.
	It requires a title and a description
	for the task.

	Syntax:
		Command-Line-Task-Manager add [title] [description]
	
	Example:
		Command-Line-Task-Manager add Complete-Project Finish the final report and submit it by Friday.


show - View Specific Task
	The show command displays the
	details of a specific task based
	on its title.

	Syntax:
		Command-Line-Task-Manager show [title]
		
	Example:
		Command-Line-Task-Manager show Complete-Project
	

done - Mark Task as Done
	The done command allows you to mark
	a task as completed.
	Specify the title of the task you
	want to mark as done.

	Syntax:
		Command-Line-Task-Manager done [title]
	
	Example:
		Command-Line-Task-Manager done Complete-Project

init - Set up the task directory.
	Example:
		Command-Line-Task-Manager init
		
settings - Update the task directory location
	This command not copy the tasks
	you already have created
	
	Syntax:
		$ Command-Line-Task-Manager settings [new-directory-path]`
)

func (s *TasksService) Initialize(cmd *cmdsModel.Command) (string, error) {
	settings, err := settings.LoadSettings()
	if err != nil {
		return errorStr, errors.New(initializeErrorStr)
	}

	path := filepath.Dir(settings.FilePath)
	err = os.MkdirAll(path, fs.ModePerm)
	if err != nil {
		return errorStr, errors.New(initializeErrorStr)
	}

	return "Command-Line-Task-Manager Initialized", nil
}

func (s *TasksService) AddTask(cmd *cmdsModel.Command) (string, error) {
	title := cmd.Args[0]
	desc := strings.Join(cmd.Args[1:], " ")
	if title == "" {
		return errorStr, errors.New("not possible to create a task with empty title")
	}
	if desc == "" {
		return errorStr, errors.New("not possible to create a task with empty description")
	}
	taskToAdd := models.Task{
		Title:       title,
		Description: desc,
		IsCompleted: false,
		CreatedAt:   time.Now(),
	}
	err := s.TasksRepo.Save(taskToAdd)
	if err != nil {
		return errorStr, errors.New("not possible to create a task")
	}
	return taskToAdd.String(), nil
}

func (s *TasksService) UpdateDescription(cmd *cmdsModel.Command) (string, error) {
	title := cmd.Args[0]
	desc := strings.Join(cmd.Args[1:], " ")

	task, err := s.TasksRepo.FindByTitle(title)
	if err != nil {
		return errorStr, errors.New("Can not find a task with title " + title)
	}
	task.ChangeDescription(desc)
	s.TasksRepo.Save(task)
	return task.String(), nil
}

func (s *TasksService) FindTask(cmd *cmdsModel.Command) (string, error) {
	task, err := s.TasksRepo.FindByTitle(cmd.Args[0])
	if err != nil {
		return errorStr, err
	}

	return task.String(), nil
}

func (s *TasksService) PrintAllTasks(_ *cmdsModel.Command) (string, error) {
	tasks, err := s.TasksRepo.FindAll()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, task := range tasks {
		fmt.Println(task)
	}
	return "", nil
}

func (s *TasksService) CompleteTask(cmd *cmdsModel.Command) (string, error) {
	title := cmd.Args[0]
	task, err := s.TasksRepo.FindByTitle(title)
	if err != nil {
		return errorStr, errors.New("Can not find a task with title " + title)
	}
	task.SetComplete()
	s.TasksRepo.Save(task)
	return task.String(), nil
}

func (s *TasksService) UpdateTasksSettings(cmd *cmdsModel.Command) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(`
Specify the absolute path where you like to save the tasks file: `)

	if scanner.Scan() {
		path := scanner.Text()
		path = filepath.Clean(path)

		if !filepath.IsAbs(path) {
			return errorStr, errors.New("file path must be absolute")
		}

		filePath := path + "/tasks.json"
		settingsToUpdate := settings.NewSettings(filePath)
		_, err := settings.UpdateSettings(*settingsToUpdate)
		if err != nil {
			return errorStr, errors.New(initializeErrorStr)
		}

		// We create the dir provided by user
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return errorStr, errors.New(initializeErrorStr)
		}
	}
	return "Command-Line-Task-Manager settings updated", nil
}

func (s *TasksService) PrintFullDocs(_ *cmdsModel.Command) (string, error) {
	fmt.Print(FULL_DOCS)

	return "", nil
}

func NewTasksService(repo *repositories.TasksRepository) *TasksService {
	return &TasksService{
		repo,
	}
}
