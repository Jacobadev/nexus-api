package testsetup

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"testing"

	"github.com/gateway-address/config"
	"github.com/gateway-address/internal/auth/repository"
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/db/postgres"
	"github.com/jmoiron/sqlx"
)

const AuthURL = "http://localhost:5000/api/v1/auth/"

func UserFromResponse(res *http.Response) (*model.User, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}
	return user, nil
}

func GetDB() (*sqlx.DB, error) {
	// Load and parse configuration
	cfgFile, err := config.LoadConfig("../../../config/config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	// Initialize and connect to the database
	db, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("NewPsqlDB: %v", err)
	}

	return db, err // Return error if user not found or any other error occurred
}

func UserIsRegistered(t *testing.T, id int) bool {
	db, err := GetDB()
	if err != nil {
		t.Errorf("GetDB: %v", err)
	}
	defer db.Close()
	repo := repository.NewRepositorySqlite(db)
	user, err := repo.GetByID(id)
	if err == nil || user != nil {
		return true
	}
	return false
}

func UserGenerator() *model.User {
	user := &model.User{}
	{
	}
	user.Email = RandomString(8) + "@godaddy.com"
	user.ID = rand.Intn(1000000)
	user.UserName = RandomString(8)
	user.Password = RandomString(8) + "!Ao"
	user.FirstName = "John"
	user.LastName = "Doe"
	return user
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// UserGenerator generates a random valid user

func SetupSuite(t *testing.T, user *model.User) {
	db, err := GetDB()
	if err != nil {
		t.Errorf("GetDB: %v", err)
	}
	defer db.Close()
	if UserIsRegistered(t, user.ID) {
		return
	}
	// Gerar token JWT para o usuário de login
	err = user.HashPassword()
	if err != nil {
		t.Errorf("HashPassword: %v", err)
	}
	_, err = db.Exec("INSERT INTO users (user_id, first_name, last_name, email, username, password) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID, user.FirstName, user.LastName, user.Email, user.UserName, user.Password)
	if err != nil {
		t.Errorf("DB Exec: %v", err)
	}
}

func LoginSetupResponse(t *testing.T) (*http.Response, *model.User, error) {
	user := UserGenerator()
	password := user.Password
	SetupSuite(t, user)
	user.Password = password

	loginJSON, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Erro ao transformar os detalhes do login em JSON: %v", err)
		return nil, nil, err
	}

	client := &http.Client{}
	res, err := client.Post(AuthURL+"login", "application/json", bytes.NewBuffer(loginJSON))
	return res, user, err
}
