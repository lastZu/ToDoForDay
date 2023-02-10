package src

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	description string
	completed   bool
}

type Operation func(todos []Task, command []string) []Task

func Initialize() map[string]Operation {
	return map[string]Operation{
		"add":    Add,
		"list":   ShowList,
		"delete": Delete,
		"done":   Done,
		"quit":   Quit,
	}
}

func Add(todos []Task, command []string) []Task {
	answer := "Please enter a task description."
	if len(command) > 1 {
		newTask := Task{
			description: strings.Join(command[1:], " "),
			completed:   false,
		}
		todos = append(todos, newTask)
		answer = "Task added."
	}
	fmt.Println(answer)
	return todos
}

func ShowList(todos []Task, command []string) []Task {
	taskView := "%d. [%s] %s\n"
	for i, todo := range todos {
		done := " "
		if todo.completed {
			done = "X"
		}
		fmt.Printf(taskView, i+1, done, todo.description)
	}
	return todos
}

func Delete(todos []Task, command []string) []Task {
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
	return todos
}

func Done(todos []Task, command []string) []Task {
	answer := "Please enter a task number."

	if len(command) > 1 {
		answer = "Invalid task number."

		index, err := strconv.Atoi(command[1])
		goodIndex := index >= 1 || index <= len(todos)
		if err == nil && goodIndex {
			answer = fmt.Sprintf("Task %d completed.", index)
			todos[index-1].completed = true
		}
	}
	fmt.Println(answer)
	return todos
}

func Quit(todos []Task, command []string) []Task {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return todos
}
