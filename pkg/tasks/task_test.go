package tasks

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TaskTestSuite struct {
	suite.Suite
}

func TestTaskSuite(t *testing.T) {
	suite.Run(t, &TaskTestSuite{})
}

func (ts *TaskTestSuite) TestTaskCompleted() {
	task := NewTask("task1", "desc")

	task.SetComplete()

	ts.Equal(task.IsCompleted, true)
}

func (ts *TaskTestSuite) TestTaskUncompleted() {
	task := NewTask("task1", "desc")

	task.SetUncomplete()

	ts.Equal(task.IsCompleted, false)
}

func (ts *TaskTestSuite) TestTaskUpdateTitle() {
	task := NewTask("task1", "desc")

	newTitle := "new title"
	task.UpdateTitle(newTitle)

	ts.Equal(task.Title, newTitle)
}

func (ts *TaskTestSuite) TestTaskUpdateDesc() {
	task := NewTask("task1", "desc")

	newDescription := "new desc"
	task.UpdateDescription(newDescription)

	ts.Equal(task.Description, newDescription)
}

func (ts *TaskTestSuite) TestTaskAppendDesc() {
	task := NewTask("task1", "desc")

	toAppend := " appended this"
	task.AppendDescription(toAppend)

	ts.Equal(task.Description, "desc appended this")
}

func (ts *TaskTestSuite) TestTaskString() {
	dateInTest, _ := time.Parse("02/01/2006", "06/07/2023")

	type fields struct {
		Title       string
		Description string
		IsCompleted bool
		CreatedAt   time.Time
		CompletedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test task uncompleted",
			fields: fields{
				Title:       "task1",
				Description: "Desc",
				IsCompleted: false,
				CreatedAt:   dateInTest,
				CompletedAt: time.Now(),
			},
			want: `
task1
Desc
Not completed
Created at 2023-07-05 19:00:00`,
		},
		{
			name: "Test task completed",
			fields: fields{
				Title:       "task1",
				Description: "Desc",
				IsCompleted: true,
				CreatedAt:   dateInTest,
				CompletedAt: dateInTest,
			},
			want: `
task1
Desc
Completed at 2023-07-05 19:00:00
Created at 2023-07-05 19:00:00`,
		},
	}
	for _, tt := range tests {
		ts.T().Run(tt.name, func(t *testing.T) {
			task := Task{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				IsCompleted: tt.fields.IsCompleted,
				CreatedAt:   tt.fields.CreatedAt,
				CompletedAt: tt.fields.CompletedAt,
			}
			if got := task.String(); got != tt.want {
				t.Errorf("Task.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
