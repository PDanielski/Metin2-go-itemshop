package main

import (
	"net/http"
)

//Middleware is a function which wraps an http.Handler and provides some additional functionality
//Can be also used for overriding completely the behavior of the http.Handler given in the input
type Middleware func(h http.Handler) http.Handler

//ChainMiddlewares is a utility function which merges a set of Middlewares and returns the merged Middleware
//The first Middleware in the set is evaluated last
func ChainMiddlewares(middlewares ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for _, middleware := range middlewares {
			h = middleware(h)
		}
		return h
	}
}
