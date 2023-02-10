package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"src.command"

	"github.com/c-bata/go-prompt"
)

var todos []command.Task

func main() {
	commands := command.Initialize()
	question := questionToUser(commands)
	for {
		result := prompt.Input(question, completer)
		fmt.Println()
		userText := scanCommand()
		userText = append(userText, result)
		command := userText[0]
		operator := commands[command]
		if operator == nil {
			fmt.Println("Unknown command. Please try again.")
			continue
		}
		todos = operator(todos, userText)
	}
}

func questionToUser(commandsList map[string]command.Operation) string {
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

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
