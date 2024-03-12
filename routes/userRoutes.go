package routes

import (
	"fmt"
	"io"
	"net/http"
)

func routeHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		userGet(w, r)
	}
	if r.Method == "POST" {
		userPost(w, r)
	}
	if r.Method == "UPDATE" {
	}
}

func userGet(w http.ResponseWriter, r *http.Request) {
	user, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s", user)
}

func userPost(w http.ResponseWriter, r *http.Request)
