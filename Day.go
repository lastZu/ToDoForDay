package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		command := strings.Fields(input.Text())

		switch command[0] {
		case "add":
			if len(command) > 1 {
				todos = append(todos, task{description: strings.Join(command[1:], " "), completed: false})
				fmt.Println("Task added.")
			} else {
				fmt.Println("Please enter a task description.")
			}
		case "list":
			for i, todo := range todos {
				if todo.completed {
					fmt.Printf("%d. [Done] %s\n", i+1, todo.description)
				} else {
					fmt.Printf("%d. %s\n", i+1, todo.description)
				}
			}
		case "delete":
			if len(command) > 1 {
				index, err := strconv.Atoi(command[1])
				if err != nil || index < 1 || index > len(todos) {
					fmt.Println("Invalid task number.")
				} else {
					todos = append(todos[:index-1], todos[index:]...)
					fmt.Println("Task deleted.")
				}
			} else {
				fmt.Println("Please enter a task number.")
			}
		case "done":
			if len(command) > 1 {
				index, err := strconv.Atoi(command[1])
				if err != nil || index < 1 || index > len(todos) {
					fmt.Println("Invalid task number.")
				} else {
					todos[index-1].completed = true
					fmt.Printf("Task %d completed.\n", index)
				}
			} else {
				fmt.Println("Please enter a task number.")
			}
		case "quit":
			fmt.Println("Goodbye!")
			os.Exit(0)
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
