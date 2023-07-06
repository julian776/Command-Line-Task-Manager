package commands

import (
	"os"
	"reflect"
	"testing"

	"github.com/julian776/Command-Line-Task-Manager/commands/models"
)

func TestBuildCommmand(t *testing.T) {
	cmdType := os.Args[1]

	tests := []struct {
		name string
		want *models.Command
	}{
		{
			name: "Should create a correct command",
			want: &models.Command{
				CmdType: cmdType,
				Args:    []string{},
				Flags: map[string]string{
					"filter": "uncompleted",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildCommmand(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildCommmand() = %v, want %v", got, tt.want)
			}
		})
	}
}
