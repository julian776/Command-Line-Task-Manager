package commands

import (
	"flag"
	"os"
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

	if len(os.Args) < 2 {
		return cmd
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
func addFlags(cmd *Command) {
	filter := flag.String("filter", "uncompleted", "Command-Line-Task-Manager ls (--filter all | uncompleted)")

	cmd.Flags["filter"] = *filter
}
