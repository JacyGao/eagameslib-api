package router

import "net/http"
import "github.com/eagames/handler"

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
		handler.Index,
	},
	Route{
		"TitleGet",
		"GET",
		"/title",
		handler.TitleGet,
	},
	Route{
		"TitlePost",
		"POST",
		"/title",
		handler.TitlePost,
	},
}
