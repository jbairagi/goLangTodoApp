package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func ShowTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := DbFindTodo(todoId)
	if todo.Id > 0 {
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := DbDestroyTodo(todoId)
	if todo.Id > 0 {
		if err := json.NewEncoder(w).Encode(todos); err != nil {
			panic(err)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Could not find Todo with this id to delete"}); err != nil {
		panic(err)
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	DbCreateTodo(todo)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}
