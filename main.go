package main

import (
	"bufio"
	"fmt"
	"os"
	"toDoList/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		command := commands.ReadCommand(scanner)
		fmt.Println(command)
	}
}
