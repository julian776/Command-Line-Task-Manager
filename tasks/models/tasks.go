package models

type Task struct {
	Title       string
	Description string
	isCompleted bool
}

func (task Task) String() string {
	taskString := "\n" + task.Title + "\n" + task.Description + "\n"
	if task.isCompleted {
		taskString += "Completed"
	} else {
		taskString += "Not completed"
	}
	return taskString
}

func (task *Task) SetComplete() {
	task.isCompleted = true
}

func (task *Task) SetUncomplete() {
	task.isCompleted = false
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
