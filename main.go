package main

import (
	"os"
	"toDoList/helpers"
)

func main() {
	router := helpers.Setup()
	router.Router(os.Args[1:])
}
