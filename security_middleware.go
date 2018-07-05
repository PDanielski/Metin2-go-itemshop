package main

import (
	"context"
	"mt2is/pkg/account"
	"net/http"

	"github.com/gorilla/sessions"
)

type key string

const accKey key = "account"

//SecurityMiddleware is a middleware used to wrap handlers which need an accountID in their context.
func SecurityMiddleware(store sessions.Store, repo account.Repository) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			sess, err := store.Get(request, "auth")
			if err != nil {
				http.Error(writer, "Problem while retrieving sessions", 500)
				return
			}

			tmpAccID, ok := sess.Values["accountID"]
			if !ok {
				http.Error(writer, "Authentication needed", 401)
				return
			}

			accountID := account.ID(tmpAccID.(int))
			acc, ok := repo.ByID(accountID)

			if !ok {
				http.Error(writer, "You are logged with an account which does not exist", 401)
				return
			}

			ctx := context.WithValue(request.Context(), accKey, acc)
			newRequest := request.WithContext(ctx)
			h.ServeHTTP(writer, newRequest)
		})
	}
}
