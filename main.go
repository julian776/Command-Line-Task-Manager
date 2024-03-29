package main

import (
	"github.com/julian776/Command-Line-Task-Manager/pkg"
	"github.com/julian776/Command-Line-Task-Manager/pkg/commands"
)

func main() {
	router := pkg.Setup()
	cmd := commands.BuildCommmand()
	router.Router(cmd)
}
