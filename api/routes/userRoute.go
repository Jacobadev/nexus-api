package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gateway-address/model"

	"github.com/gateway-address/config"
)

func ValidateUserInput(r *http.Request) error {
	return nil
}

func WriteUserGETResponse(w http.ResponseWriter, data interface{}) {
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

type userFormData struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func getResponseData(user model.User) userFormData {
	userFormData := struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}{
		UserName: user.UserName,
		Password: user.Password,
	}
	return userFormData
}

func WriteUserPOSTResponse(w http.ResponseWriter, userFormData userFormData) {
	// Create a custom JSON object containing only the UserName and Password fields

	// Marshal the custom JSON object to JSON
	jsonData, err := json.Marshal(userFormData)
	if err != nil {
		http.Error(w, "Failed to marshal JSON data", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the response writer
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Failed to write JSON response", http.StatusInternalServerError)
		return
	}
}

func UserMethodController(w http.ResponseWriter, r *http.Request) {
	if err := ValidateUserInput(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if r.Method == "GET" {
		var response interface{} = UserGetAll()
		WriteUserGETResponse(w, response)
	}
	if r.Method == "POST" {
		err := userCreate(w, r)
		if err != nil {
			http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with success message
		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("User created successfully"))
		if err != nil {
			// Failed to write response, log the error
			fmt.Println("Failed to write response:", err)
		}
	}

	// if r.Method == "UPDATE" {
	//    userUpdate()
	// }
	// if r.Method == "PATCH" {
	//    userPartialUpdate()
	// }
	// if r.Method == "DELETE" {
	//    userDelete()
	// }
}

func UserGetAll() []model.User {
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

func userCreate(w http.ResponseWriter, r *http.Request) error {
	// Declare a variable to hold the decoded data
	var credentials userFormData

	// Decode JSON request body into UserCredentials struct
	err := json.NewDecoder(r.Body).Decode(&credentials)
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
	stmt, err := db.Prepare("INSERT INTO user (user_name, password) VALUES (?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare SQL statement", http.StatusInternalServerError)
		return err
	}
	fmt.Printf("user_name: %s, password: %s", credentials.UserName, credentials.Password)
	defer stmt.Close()

	// Execute SQL statement to insert data
	_, err = stmt.Exec(credentials.UserName, credentials.Password)
	if err != nil {
		http.Error(w, "Failed to execute SQL statement", http.StatusInternalServerError)
		return err
	}

	// Return nil if the operation is successful
	return nil
}
