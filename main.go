package main

import (
	"os"

	"github.com/julian776/Command-Line-Task-Manager/helpers"
)

func main() {
	router := helpers.Setup()
	router.Router(os.Args[1:])
}
