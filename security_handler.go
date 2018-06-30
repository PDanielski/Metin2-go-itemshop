package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

type SecurityHandler struct {
	sessStore sessions.Store
	secret    string
	next      *http.Handler
}

func NewSecurityHandler(store sessions.Store, secret string, next *http.Handler) *SecurityHandler {
	return &SecurityHandler{store, secret, next}
}

func (h *SecurityHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	sess, err := h.sessStore.Get(request, "auth")
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	tmpAccountID, ok := sess.Values["accountId"]
	if !ok {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	accountID := tmpAccountID.(int)

	fmt.Fprintf(writer, "Your account id is %d", accountID)

}
