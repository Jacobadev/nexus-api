package http

import (
	"encoding/json"
	"net/http"

	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/auth"
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/httpErrors"
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
		defer r.Body.Close()
		if err := utils.ReadRequest(r, user); err != nil {
			h.logger.Infof("reading request: %s, err: %v", http.StatusText(http.StatusBadRequest), err)
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		if err := utils.ValidateUser(user); err != nil {
			h.logger.Errorf("Invalid user: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		createdUser, registerErr := h.authUC.Register(user)
		if registerErr != nil {
			httpErrors.WriteJsonResponse(w, registerErr)
			return
		}
		h.logger.Infof("User created ID: %d", createdUser.User.ID)
		userJson, err := json.Marshal(createdUser)
		if err != nil {
			h.logger.Errorf("Error creating user: %v", err)
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(userJson)
	}
}

func (h *authHandlers) Login() http.HandlerFunc {
	type Login struct {
		UserName string `json:"username" db:"username"`
		Email    string `json:"email" db:"email" validate:"omitempty,lte=60,email"`
		Password string `json:"password,omitempty" db:"password" validate:"required,gte=6"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the JSON body into a Login struct
		login := &Login{}
		if err := utils.ReadRequest(r, login); err != nil {
			h.logger.Infof("err reading request: %s, err: %v", http.StatusText(http.StatusBadRequest), err)
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		authenticatedUser, err := h.authUC.Login(&model.User{
			UserName: login.UserName,
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			h.logger.Errorf("Authentication failed: %v", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

		}
		authdUserJson, err := json.Marshal(authenticatedUser)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		// Write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.Write(authdUserJson)
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

//	func (h *authHandlers) PartialUpdate() http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			vars := mux.Vars(r)
//			idStr := vars["id"]
//			id, err := strconv.Atoi(idStr)
//			if err != nil {
//				http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
//				return
//			}
//			var user user.User
//			err = json.NewDecoder(r.Body).Decode(&user)
//			if err != nil {
//				http.Error(w, fmt.Sprintf("Failed to convert id to int: %v", err), http.StatusInternalServerError)
//				return
//			}
//			err = h.authUC.PartialUpdateByID(id, &user)
//			if err != nil {
//				http.Error(w, fmt.Sprintf("Failed to update user: %v", err), http.StatusInternalServerError)
//			}
//		}
//	}
