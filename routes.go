package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetAllTodo",
		"GET",
		"/todos",
		GetAllTodo,
	},
	Route{
		"CreateTodo",
		"POST",
		"/todos",
		CreateTodo,
	},
	Route{
		"ShowTodo",
		"GET",
		"/todos/{todoId}",
		ShowTodo,
	},
	Route{
		"DeleteTodo",
		"DELETE",
		"/deleteTodo/{todoId}",
		DeleteTodo,
	},
}
