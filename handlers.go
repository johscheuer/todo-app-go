package main

import (
	//	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "A simple todo App :)")
}

func ShowAllTodo(w http.ResponseWriter, r *http.Request) {
	RedisGetAllTodos()
	fmt.Fprintln(w, "Show all todos")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	RedisDeleteTodo(name)
	fmt.Fprintln(w, "delete a todo name %s", name)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	RedisCreateTodo(name)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "create a todo name %s", name)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "alive")
}
