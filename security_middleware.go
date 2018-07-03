package main

import (
	"context"
	"net/http"

	"github.com/gorilla/sessions"
)

type key string

const accIDKey key = "accountID"

//SecurityMiddleware is a middleware used to wrap handlers which need an accountID in their context.
func SecurityMiddleware(store sessions.Store) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			sess, err := store.Get(request, "auth")
			if err != nil {
				http.Error(writer, "Problem while retrieving sessions", 500)
				return
			}

			accountID, ok := sess.Values["accountID"]
			if !ok {
				http.Error(writer, "Authentication needed", 401)
				return
			}

			ctx := context.WithValue(request.Context(), accIDKey, accountID.(int))
			newRequest := request.WithContext(ctx)
			h.ServeHTTP(writer, newRequest)
		})
	}
}
