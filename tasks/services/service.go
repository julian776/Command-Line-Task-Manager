package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/julian776/Command-Line-Task-Manager/tasks/models"
	"github.com/julian776/Command-Line-Task-Manager/tasks/repositories"
	"github.com/julian776/Command-Line-Task-Manager/tasks/settings"
)

type TasksService struct {
	settings settings.Settings
}

const (
	errorStr = "ERROR"
)

func (s *TasksService) AddTask(params []string) (string, error) {
	title := params[0]
	desc := strings.Join(params[1:], " ")
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
	repositories.Save(s.settings.FileName, taskToAdd)
	return "Task created", nil
}

func (s *TasksService) UpdateDescription(params []string) (string, error) {
	title := params[0]
	desc := strings.Join(params[1:], " ")
	task, err := repositories.FindByTitle(s.settings.FileName, title)
	if err != nil {
		task.ChangeDescription(desc)
		repositories.Save(s.settings.FileName, task)
		return "Task Updated", nil
	} else {
		return errorStr, errors.New("Can not find a task with title " + title)
	}
}

func (s *TasksService) FindTask(params []string) (string, error) {
	task, err := repositories.FindByTitle(s.settings.FileName, params[0])
	if err != nil {
		return "", err
	} else {
		return task.String(), nil
	}
}

func (s *TasksService) PrintAllTasks(_ []string) (string, error) {
	tasks, err := repositories.FindAll(s.settings.FileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, task := range tasks {
		fmt.Println(task)
	}
	return "", nil
}

func (s *TasksService) CompleteTask(params []string) (string, error) {
	title := params[0]
	task, err := repositories.FindByTitle(s.settings.FileName, title)
	if err != nil {
		return errorStr, errors.New("Can not find a task with title " + title)
	}
	task.SetComplete()
	repositories.Save(s.settings.FileName, task)
	return "Task Completed", nil
}

func (s *TasksService) PrintFullDocs(params []string) (string, error) {
	fmt.Println(`
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
	`)

	return "", nil
}

func (s *TasksService) NewTasksService(settings settings.Settings) {
	s.settings = settings
}
