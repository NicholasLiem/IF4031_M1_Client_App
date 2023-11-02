package structs

import "net/http"

/**
Route Struct
Source code: https://github.com/hellojebus/go-mux-jwt-boilerplate
*/

var AppRoutes []RoutePrefix

type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}
