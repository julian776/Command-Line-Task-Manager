package repositories

import (
	"os"
	"testing"

	"github.com/julian776/Command-Line-Task-Manager/tasks/models"
	"github.com/stretchr/testify/suite"
)

const (
	FILE_PATH = "test_file.json"
)

type TaskRepositoryTests struct {
	suite.Suite
	filePath string
}

func TestRepoTasks(t *testing.T) {
	suite.Run(t, &TaskRepositoryTests{
		filePath: FILE_PATH,
	})
}

func (trs *TaskRepositoryTests) TestSaveTask() {
	err := Save(trs.filePath, models.Task{})
	trs.Nil(err)
}

func (trs *TaskRepositoryTests) TestFindAll() {
	mapTasks, err := FindAll(trs.filePath)
	trs.Nil(err)
	trs.Len(mapTasks, 0)
}

func (trs *TaskRepositoryTests) TestFindAllErrorNoFile() {
	wrongFilePath := "fghfghfghgf.json"
	//emptyTask := models.Task{}

	mapTasks, err := FindAll(wrongFilePath)
	trs.Error(err)
	trs.Len(mapTasks, 0)
}

func (trs *TaskRepositoryTests) TestFindByTitle() {
	taskTitle := "task1"
	task := models.Task{
		Title: taskTitle,
	}
	Save(trs.filePath, task)

	taskGot, err := FindByTitle(trs.filePath, taskTitle)
	trs.Nil(err)
	trs.Equal(task, taskGot)
}

func (trs *TaskRepositoryTests) TestFindByTitleError() {
	taskTitle := "task1"
	emptyTask := models.Task{}

	taskGot, err := FindByTitle(trs.filePath, taskTitle)
	trs.Error(err)
	trs.Equal(emptyTask, taskGot)
}

func (trs *TaskRepositoryTests) SetupSuite() {
	file, err := os.Create(trs.filePath)
	if err != nil {
		trs.FailNowf("can not setup file test", err.Error())
	}
	_, err = file.Write([]byte("{}"))
	if err != nil {
		trs.FailNowf("can not setup file test", err.Error())
	}
}

func (trs *TaskRepositoryTests) TearDownTest() {
	err := os.WriteFile(trs.filePath, []byte("{}"), os.ModeAppend)
	if err != nil {
		trs.FailNowf("can clean file test", err.Error())
	}
}

func (trs *TaskRepositoryTests) TearDownSuite() {
	err := os.Remove(trs.filePath)
	if err != nil {
		trs.T().Error("can not clean test_file.json")
	}
}
