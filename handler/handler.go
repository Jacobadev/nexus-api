package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gateway-address/repository"
	"github.com/gateway-address/user"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserRepository user.UserRepository
	Repository     *repository.RepositorySqlite
}

func NewUserHandler(userRepository user.UserRepository, repo *repository.RepositorySqlite) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
		Repository:     repo,
	}
}

func GetUsersHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userHandler.Repository.GetAll()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get users: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUserHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validate.ValidateRequest(r)
		user, err := user.ExtractUserInput(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = userHandler.Repository.Create(user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create user: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetUserByIDHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
			return
		}

		user, err := userHandler.Repository.GetByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)

		// Use the 'user' object as needed
	}
}

func getParamsFromRequest(r *http.Request) (int, int, error) {
	const maxLimit = 20
	limit := mux.Vars(r)["limit"] // Access captured limit
	offset := mux.Vars(r)["offset"]

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = maxLimit // Set default value to maxLimit if limit is invalid or not provided
	} else if limitInt > maxLimit {
		limitInt = maxLimit // Cap limit to maxLimit
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return 0, 0, err
	}

	return limitInt, offsetInt, nil
}

func GetPaginatedUsersHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset, err := getParamsFromRequest(r)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get params: %v", err), http.StatusInternalServerError)
			return
		}

		users, err := userHandler.Repository.GetPaginated(limit, offset)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get paginated users: %v", err), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func DeleteUserByIDHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
			return
		}
		err = userHandler.Repository.DeleteByID(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete user: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateUserByIDHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
			return
		}
		var user user.User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
			return
		}
		err = userHandler.Repository.UpdateByID(id, &user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to update user: %v", err), http.StatusInternalServerError)
		}
	}
}

func PartialUpdateUserByIDHandler(userHandler *UserHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
			return
		}
		var user user.User
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
			return
		}
		err = userHandler.Repository.PartialUpdateByID(id, &user)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to update user: %v", err), http.StatusInternalServerError)
		}
	}
}
