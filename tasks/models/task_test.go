package models

import (
	"testing"
	"time"
)

func TestTask_String(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
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
