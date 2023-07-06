package commands

import (
	"flag"
	"os"

	"github.com/julian776/Command-Line-Task-Manager/commands/models"
)

func BuildCommmand() *models.Command {
	cmd := &models.Command{
		Flags: map[string]string{},
	}

	addFlags(cmd)
	// Must be called after addFlags do not move
	flag.Parse()

	cmd.CmdType = os.Args[1]
	cmd.Args = flag.Args()

	if len(cmd.Args) > 1 {
		cmd.Args = cmd.Args[1:]
	}

	return cmd
}

// Adds the flags to the command object.
// By the moment this is done manually.
func addFlags(cmd *models.Command) {
	filter := flag.String("filter", "uncompleted", "Command-Line-Task-Manager ls (--filter all | uncompleted)")

	cmd.Flags["filter"] = *filter
}
