package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gateway-address/config"
	"github.com/gateway-address/model"
)

func UserGetAll() []model.User {
	var users []model.User

	db, err := config.GetDbConnection()
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	defer db.Close()

	allUsers, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer allUsers.Close()

	for allUsers.Next() {
		var user model.User
		if err := allUsers.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}
	if err := allUsers.Err(); err != nil {
		fmt.Println(err)
		return nil
	}
	return users
}

func UserCreate(w http.ResponseWriter, r *http.Request) error {
	var userInfo model.User
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, "Failed to decode JSON data", http.StatusBadRequest)
		return err
	}

	// Get database connection
	db, err := config.GetDbConnection()
	if err != nil {
		http.Error(w, "Failed to connect to the database", http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	// Prepare SQL statement
	stmt, err := db.Prepare("INSERT INTO user (first_name,last_name,user_name, password,email,) VALUES (?, ?,?,?,?)")
	if err != nil {
		http.Error(w, "Failed to prepare SQL statement", http.StatusInternalServerError)
		return err
	}
	defer stmt.Close()

	// Execute SQL statement to insert data
	_, err = stmt.Exec(userInfo.FirstName, userInfo.LastName, userInfo.UserName, userInfo.Password, userInfo.Email)
	if err != nil {
		http.Error(w, "Failed to execute SQL statement", http.StatusInternalServerError)
		return err
	}

	return nil
}
