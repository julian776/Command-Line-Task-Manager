package main

import (
	"bufio"
	"os"
	"toDoList/helpers"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	router := helpers.Setup()
	for true {
		command := helpers.ReadCommand(scanner)
		router.Router(command)
	}
}
