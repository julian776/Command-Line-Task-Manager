package models

import (
	"fmt"
	"time"
)

type Task struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	IsCompleted bool      `json:"isCompleted,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	CompletedAt time.Time `json:"completedAt,omitempty"`
}

func (task Task) String() string {
	taskString := "\n" + task.Title + "\n" + task.Description + "\n"
	if task.IsCompleted {
		taskString += fmt.Sprintf("Completed at %s", task.CompletedAt.Local().Format(time.DateTime))
	} else {
		taskString += "Not completed"
	}
	taskString += fmt.Sprintf("\nCreated at %s", task.CreatedAt.Local().Format(time.DateTime))
	return taskString
}

func (task *Task) SetComplete() {
	task.IsCompleted = true
	task.CompletedAt = time.Now()
}

func (task *Task) SetUncomplete() {
	task.IsCompleted = false
}

func (task *Task) ChangeDescription(newDescription string) string {
	task.Description = newDescription
	return "Description updated"
}

func (task *Task) AppendDescription(text string) string {
	task.Description = task.Description + text
	return task.Description
}

func (task *Task) UpdateTitle(newTitle string) string {
	task.Title = newTitle
	return "Title updated"
}
