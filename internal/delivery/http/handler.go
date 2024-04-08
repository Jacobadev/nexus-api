package http

import (
	"fmt"
	"net/http"

	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/auth"
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/logger"
	"github.com/gateway-address/pkg/utils"
)

type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	logger logger.Logger
}

func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase, log logger.Logger) *authHandlers {
	return &authHandlers{cfg: cfg, authUC: authUC, logger: log}
}

func (h *authHandlers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}

		if err := utils.ReadRequest(r, user); err != nil {
			utils.LogResponseError(r, h.logger, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(user)
		createdUser, err := h.authUC.Register(user)
		if err != nil {
			utils.LogResponseError(r, h.logger, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(createdUser)

		// fmt.Println("creating")
		//
		// if err := utils.ReadRequest(r, user); err != nil {
		// 	utils.LogResponseError(r, h.logger, err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		// w.WriteHeader(http.StatusCreated)
		// jsonData, err := json.Marshal(createdUser)
		// if err != nil {
		// 	utils.LogResponseError(r, h.logger, err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		//
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusCreated)
		// w.Write(jsonData)
	}
}

//
// func (h *authHandlers) GetUsers() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		users, err := h.authUC.GetAll()
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to get users: %v", err), http.StatusInternalServerError)
// 			return
// 		}
//
// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(users)
// 	}
// }
//
// func (h *authHandlers) GetUserByID() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		idStr := vars["id"]
//
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
// 			return
// 		}
//
// 		user, err := h.authUC.GetByID(id)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
// 			return
// 		}
//
// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(user)
//
// 		// Use the 'user' object as needed
// 	}
// }
//
// func (h *authHandlers) Delete() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		idStr := vars["id"]
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		err = h.authUC.Delete(id)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to delete user: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 	}
// }
//
// func (h *authHandlers) UpdateByID() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		idStr := vars["id"]
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		var user user.User
// 		err = json.NewDecoder(r.Body).Decode(&user)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		err = h.authUC.UpdateByID(id, &user)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to update user: %v", err), http.StatusInternalServerError)
// 		}
// 	}
// }

// func (h *authHandlers) PartialUpdate() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		idStr := vars["id"]
// 		id, err := strconv.Atoi(idStr)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		var user user.User
// 		err = json.NewDecoder(r.Body).Decode(&user)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		err = h.authUC.PartialUpdateByID(id, &user)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Failed to update user: %v", err), http.StatusInternalServerError)
// 		}
// 	}
// }
