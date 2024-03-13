package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gateway-address/api/model"

	"github.com/gateway-address/api/config"

	_ "github.com/mattn/go-sqlite3"
)

func ValidateUserInput(r *http.Request) error {
	return nil
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}) {
	users, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to unmarshal JSON data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(users)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func UserMethodController(w http.ResponseWriter, r *http.Request) {
	if err := ValidateUserInput(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var response interface{}

	if r.Method == "GET" {
		response = UserGetAll(r, w)

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
	WriteJSONResponse(w, response)
}

func UserGetAll(r *http.Request, w http.ResponseWriter) []model.User {
	var users []model.User

	db, err := config.GetDbConnection()
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil
	}
	return users
}
