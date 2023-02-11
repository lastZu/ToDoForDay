package main

import (
	"ToDoDay/src"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todos []src.Task

func main() {
	commands := src.Initialize()
	question := questionToUser(commands)
	for {
		fmt.Println(question)
		userText := scanCommand()
		command := userText[0]
		operator := commands[command]
		if operator == nil {
			fmt.Println("Unknown command. Please try again.")
			continue
		}
		todos = operator(todos, userText)
	}
}

func questionToUser(commandsList map[string]src.Operation) string {
	var names []string
	for name := range commandsList {
		names = append(names, name)
	}
	separateNames := strings.Join(names, ", ")
	enter := "Enter a command (%s):"
	return fmt.Sprintf(enter, separateNames)
}

func scanCommand() []string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	command := strings.Fields(input.Text())
	return command
}
