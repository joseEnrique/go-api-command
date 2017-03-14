package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var initAPI = "/api/v1"

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		Index,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		Index,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		Index,
	},
	Route{
		"PostCommand",
		"POST",
		initAPI + "/command/create",
		CommandExecuted,
	},
}
