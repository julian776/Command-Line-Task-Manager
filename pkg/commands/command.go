package commands

import (
	"flag"
	"os"
)

var (
	filter = flag.String("filter", "uncompleted", "Command-Line-Task-Manager ls (--filter all | uncompleted)")
)

type Command struct {
	CmdType string
	Args    []string
	Flags   map[string]string
}

func BuildCommmand() *Command {
	cmd := &Command{
		Flags: map[string]string{},
	}

	addFlags(cmd)

	if len(os.Args) < 2 {
		return cmd
	}

	cmd.Args = flag.Args()[1:]
	cmd.CmdType = flag.Args()[0]

	if len(cmd.Args) > 1 {
		cmd.Args = cmd.Args[1:]
	}

	return cmd
}

// Adds the flags to the command object.
// By the moment this is done manually.
func addFlags(cmd *Command) {
	flag.Parse()
	cmd.Flags["filter"] = *filter
}
