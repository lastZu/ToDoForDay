package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type task struct {
	description string
	completed   bool
}

var todos []task
var commands []string

func main() {
	commands = append(commands, "add", "list", "delete", "done", "quit")
	for {
		askComand()
		command := scanCommand()

		switch command[0] {
		case "add":
			todos = add(todos, command)
		case "list":
			showList(todos)
		case "delete":
			todos = delete(todos, command)
		case "done":
			done(todos, command)
		case "quit":
			quit()
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}

func askComand() {
	enter := "Enter a command (%s):\n"
	allCommand := strings.Join(commands, ", ")
	fmt.Printf(enter, allCommand)
}

func scanCommand() []string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	command := strings.Fields(input.Text())
	return command
}
