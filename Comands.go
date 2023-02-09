package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type task struct {
	description string
	completed   bool
}

type operation func(todos []task, command []string) []task

func commandInitialize() map[string]operation {
	return map[string]operation{
		"add":    add,
		"list":   showList,
		"delete": delete,
		"done":   done,
		"quit":   quit,
	}
}

func add(todos []task, command []string) []task {
	ansver := "Please enter a task description."
	if len(command) > 1 {
		newTask := task{
			description: strings.Join(command[1:], " "),
			completed:   false,
		}
		todos = append(todos, newTask)
		ansver = "Task added."
	}
	fmt.Println(ansver)
	return todos
}

func showList(todos []task, command []string) []task {
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

func delete(todos []task, command []string) []task {
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

func done(todos []task, command []string) []task {
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

func quit(todos []task, command []string) []task {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return todos
}
