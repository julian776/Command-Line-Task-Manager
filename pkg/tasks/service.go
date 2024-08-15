package tasks

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/julian776/Command-Line-Task-Manager/pkg/commands"
)

type TasksService struct {
	TasksRepo *TasksRepository
}

const (
	errorStr = "ERROR"

	initializeErrorStr = "not posible to initialize Command-Line-Task-Manager, you are using root user?"

	fullDocs = `
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
		Command-Line-Task-Manager init`
)

func (s *TasksService) Initialize(cmd *commands.Command) (string, error) {
	err := CreateDefaultSettingsFile()
	if err != nil {
		return errorStr, errors.New(initializeErrorStr)
	}

	return "Command-Line-Task-Manager Initialized", nil
}

func (s *TasksService) AddTask(cmd *commands.Command) (string, error) {
	title := cmd.Args[0]
	desc := strings.Join(cmd.Args[1:], " ")
	if title == "" {
		return errorStr, errors.New("not possible to create a task with empty title")
	}
	if desc == "" {
		return errorStr, errors.New("not possible to create a task with empty description")
	}
	taskToAdd := Task{
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

func (s *TasksService) UpdateDescription(cmd *commands.Command) (string, error) {
	title := cmd.Args[0]
	desc := strings.Join(cmd.Args[1:], " ")

	task, err := s.TasksRepo.FindByTitle(title)
	if err != nil {
		return errorStr, errors.New("Can not find a task with title " + title)
	}
	task.UpdateDescription(desc)
	s.TasksRepo.Save(task)
	return task.String(), nil
}

func (s *TasksService) FindTask(cmd *commands.Command) (string, error) {
	task, err := s.TasksRepo.FindByTitle(cmd.Args[0])
	if err != nil {
		return errorStr, err
	}

	return task.String(), nil
}

func (s *TasksService) PrintAllTasks(cmd *commands.Command) (string, error) {
	tasks, err := s.TasksRepo.FindAll()
	if err != nil {
		return "", err
	}
	for _, task := range tasks {
		if cmd.Flags["filter"] == "uncompleted" && task.IsCompleted {
			continue
		}

		fmt.Println(task)
	}
	return "", nil
}

func (s *TasksService) CompleteTask(cmd *commands.Command) (string, error) {
	title := cmd.Args[0]
	task, err := s.TasksRepo.FindByTitle(title)
	if err != nil {
		return errorStr, errors.New("Can not find a task with title " + title)
	}
	task.SetComplete()
	s.TasksRepo.Save(task)
	return task.String(), nil
}

func (s *TasksService) PrintFullDocs(_ *commands.Command) (string, error) {
	fmt.Print(fullDocs)

	return "", nil
}

func NewTasksService(repo *TasksRepository) *TasksService {
	return &TasksService{
		repo,
	}
}
