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
	if len(command) > 1 {
		todos = append(todos, task{description: strings.Join(command[1:], " "), completed: false})
		fmt.Println("Task added.")
	} else {
		fmt.Println("Please enter a task description.")
	}
	return todos
}

func showList(todos []task, command []string) []task {
	taskView := "%d. [%s] %s\n"
	for i, todo := range todos {
		done := ""
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
	return todos
}

func quit(todos []task, command []string) []task {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return todos
}
