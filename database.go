package main

import (
	"time"
)
var currentId int

var todos Todos

func init() {
	DbCreateTodo(Todo{Name: "Learn Go"})
	DbCreateTodo(Todo{Name: "Apply Go"})
}

func DbFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func DbCreateTodo(t Todo) {
	currentId += 1
	t.Id = currentId
	t.Completed = false
	t.BeginningDate = time.Now().Format("01-02-2006")
	todos = append(todos, t)
	return
}

func DbDestroyTodo(id int) Todo {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return t
		}
	}
	return Todo{}
}
