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
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ShowAllTodo",
		"GET",
		"/todo",
		ShowAllTodo,
	},
	Route{
		"DeleteTodo",
		"DELETE",
		"/todo/{name}",
		DeleteTodo,
	},
	Route{
		"CreateTodo",
		"POST",
		"/todo/{name}",
		CreateTodo,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}
