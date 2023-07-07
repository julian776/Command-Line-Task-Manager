package services

import (
	"os"
	"regexp"
	"testing"

	"github.com/julian776/Command-Line-Task-Manager/commands/models"
	"github.com/julian776/Command-Line-Task-Manager/tasks/repositories"
	"github.com/julian776/Command-Line-Task-Manager/tasks/repositories/settings"
	"github.com/stretchr/testify/suite"
)

const (
	FILE_PATH = "test_file.json"
)

type TaskServiceTests struct {
	suite.Suite
	filePath     string
	tasksService *TasksService
}

func TestServiceTasks(t *testing.T) {
	suite.Run(t, &TaskServiceTests{
		filePath:     FILE_PATH,
		tasksService: &TasksService{},
	})
}

func (trs *TaskServiceTests) TestAddTask() {
	cmd := &models.Command{
		CmdType: "add",
		Args:    []string{"titleTask", "Desc"},
	}
	strGot, err := trs.tasksService.AddTask(cmd)
	trs.Nil(err)
	trs.Regexp(regexp.MustCompile("^\ntitleTask\nDesc\nNot completed"), strGot)
}

func (trs *TaskServiceTests) TestAddTask_InvalidTitle() {
	cmd := &models.Command{
		CmdType: "add",
		Args:    []string{"", "Desc"},
	}
	strGot, err := trs.tasksService.AddTask(cmd)
	trs.Error(err)
	trs.Equal("ERROR", strGot)
}

func (trs *TaskServiceTests) TestAddTask_InvalidDescription() {
	cmd := &models.Command{
		CmdType: "add",
		Args:    []string{"title1"},
	}
	strGot, err := trs.tasksService.AddTask(cmd)
	trs.Error(err)
	trs.Equal("ERROR", strGot)
}

func (trs *TaskServiceTests) TestFindTask() {
	cmdTaskToAdd := &models.Command{
		CmdType: "add",
		Args:    []string{"task1", "desc"},
	}
	trs.tasksService.AddTask(cmdTaskToAdd)

	cmd := &models.Command{
		CmdType: "show",
		Args:    []string{"task1"},
	}
	strGot, err := trs.tasksService.FindTask(cmd)
	trs.Nil(err)
	trs.Regexp(regexp.MustCompile("^\ntask1\ndesc\nNot completed"), strGot)
}

func (trs *TaskServiceTests) TestUpdateDescription() {
	cmdTaskToAdd := &models.Command{
		CmdType: "add",
		Args:    []string{"task1", "desc"},
	}
	trs.tasksService.AddTask(cmdTaskToAdd)

	cmd := &models.Command{
		CmdType: "show",
		Args:    []string{"task1", "new", "desc"},
	}
	strGot, err := trs.tasksService.UpdateDescription(cmd)
	trs.Nil(err)
	trs.Regexp(regexp.MustCompile("^\ntask1\nnew desc\nNot completed"), strGot)
}

func (trs *TaskServiceTests) TestCompleteTask() {
	cmdTaskToAdd := &models.Command{
		CmdType: "add",
		Args:    []string{"task1", "desc"},
	}
	trs.tasksService.AddTask(cmdTaskToAdd)

	cmd := &models.Command{
		CmdType: "done",
		Args:    []string{"task1"},
	}
	strGot, err := trs.tasksService.CompleteTask(cmd)
	trs.Nil(err)
	trs.Regexp(regexp.MustCompile("^\ntask1\ndesc\nCompleted at"), strGot)
}

func (trs *TaskServiceTests) TestFindTaskErrorNoTask() {
	cmd := &models.Command{
		CmdType: "show",
		Args:    []string{"task1"},
	}
	strGot, err := trs.tasksService.FindTask(cmd)
	trs.Error(err)
	trs.Equal("ERROR", strGot)
}

func (trs *TaskServiceTests) TestLs() {
	cmd := &models.Command{
		CmdType: "ls",
	}
	trs.tasksService.PrintAllTasks(cmd)
}

func (trs *TaskServiceTests) TestHelp() {
	cmd := &models.Command{
		CmdType: "help",
	}
	strGot, err := trs.tasksService.PrintFullDocs(cmd)
	trs.Nil(err)
	trs.Equal("", strGot)
}

func (trs *TaskServiceTests) SetupTest() {
	settings := settings.Settings{
		FilePath: FILE_PATH,
	}
	// Setup Tasks Repo
	repo := repositories.NewTasksRepository(settings)

	// Setup Tasks Service
	trs.tasksService = NewTasksService(repo)
}

func (trs *TaskServiceTests) SetupSuite() {
	file, err := os.Create(trs.filePath)
	if err != nil {
		trs.FailNowf("can not setup file test", err.Error())
	}
	_, err = file.Write([]byte("{}"))
	if err != nil {
		trs.FailNowf("can not setup file test", err.Error())
	}
}

func (trs *TaskServiceTests) TearDownTest() {
	err := os.WriteFile(trs.filePath, []byte("{}"), os.ModeAppend)
	if err != nil {
		trs.FailNowf("can not clean file test", err.Error())
	}
}

func (trs *TaskServiceTests) TearDownSuite() {
	err := os.Remove(trs.filePath)
	if err != nil {
		trs.T().Error("can not clean test_file.json")
	}
}
