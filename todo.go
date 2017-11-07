package main

type Todo struct {
	Id        		int     `json:"id"`
	Name      		string  `json:"name"`
	Completed 		bool    `json:"completed"`
	BeginningDate string 	`json:"beginningDate"`
}

type Todos []Todo
