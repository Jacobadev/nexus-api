package main

import (
	"fmt"
	"net/http"

	_ "github.com/JCPURGER/gateway-address"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/user", routes.routeHandle)
	start()
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func start() {
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
