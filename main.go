package main

import (
	"github.com/julian776/Command-Line-Task-Manager/commands"
	"github.com/julian776/Command-Line-Task-Manager/helpers"
)

func main() {
	router := helpers.Setup()
	cmd := commands.BuildCommmand()
	router.Router(cmd)
}
