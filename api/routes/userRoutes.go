package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gateway-address/api/model"

	_ "github.com/mattn/go-sqlite3"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "main.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ValidateUserInput(r *http.Request) error {
	return nil
}

func UserMethodController(w http.ResponseWriter, r *http.Request) {
	if err := ValidateUserInput(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if r.Method == "GET" {
		UserGetAll(w, r)
		// TODO
		// if r.Method == "POST" {
		// 	userCreate(w, r)
		// }
		// if r.Method == "UPDATE" {
		//    userUpdate()
		// }
		// if r.Method == "PATCH" {
		//    userPartialUpdate()
		// }
		// if r.Method == "DELETE" {
		//    userDelete()
		// }
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UserGetAll(w http.ResponseWriter, r *http.Request) model.User {
	var user model.User
	db, err := connectDB()
	if err != nil {
		fmt.Print(err)
	}
	row := db.QueryRow("SELECT first_name,last_name FROM user")

	err = row.Scan(&user.FirstName, &user.LastName)
	if err != nil {
		fmt.Print(err)
	}
	return user
}

func userPost(w http.ResponseWriter, r *http.Request) {
}
