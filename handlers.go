package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "A simple todo App :)")
}

func ShowAllTodo(w http.ResponseWriter, r *http.Request) {
	todos := RedisGetAllTodos()
	jsonTodos, _ := json.Marshal(todos)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonTodos))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	RedisDeleteTodo(name)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "delete a todo name %s", name)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	RedisCreateTodo(name)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "create a todo name %s", name)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "alive")
}
