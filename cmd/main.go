package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gateway-address/handler"
	"github.com/gateway-address/repository"
	"github.com/gateway-address/server"
	"github.com/gateway-address/user"
	"github.com/gateway-address/websocket"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func RegisterUserHandler(router *mux.Router, userHandler *handler.UserHandler) {
	router.HandleFunc("/", handler.CreateUserHandler(userHandler)).Methods("POST")
	router.HandleFunc("/", handler.GetUsersHandler(userHandler)).Methods("GET")
	router.HandleFunc("/{id}", handler.GetUserByIDHandler(userHandler)).Methods("GET")
	router.HandleFunc("/{id}", handler.DeleteUserByIDHandler(userHandler)).Methods("DELETE")
	router.HandleFunc("/{id}", handler.UpdateUserByIDHandler(userHandler)).Methods("PUT")
	router.HandleFunc("/{id}", handler.PartialUpdateUserByIDHandler(userHandler)).Methods("PATCH")
}

func RegisterWebSocket(router *mux.Router) {
	router.HandleFunc("/ws", websocket.WebSocket)
}

func RootHandler(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "OK"}`)
	}).Methods("GET")
}

var ctx = context.Background()

func main() {
	r := mux.NewRouter()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	go func() {
		res := rdb.Ping(ctx)

		if err := res.Err(); err != nil {
			fmt.Println("Erro ao conectar ao Redis:", err)
		} else {
			fmt.Println("Conexão com o Redis estabelecida com sucesso!")
		}
	}()

	// ReadHeaderTimeout: 3 * time.Second,
	RegisterWebSocket(r)
	fmt.Println(os.Getenv("APPID"))
	userRouter := r.PathPrefix("/api/v1/user").Subrouter().UseEncodedPath()
	var userRepository user.UserRepository
	repo, err := repository.NewRepositorySqlite()
	if err != nil {
		fmt.Println(err)
	}
	userHandler := handler.NewUserHandler(userRepository, repo)

	RegisterUserHandler(userRouter, userHandler)

	RootHandler(r)

	server.StartServer(r)
}
