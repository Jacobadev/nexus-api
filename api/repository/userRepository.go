package repository

import (
	"database/sql"
	"fmt"

	"github.com/gateway-address/model"
	_ "modernc.org/sqlite"
)

type UserRepository interface {
	GetAll() ([]interface{}, error)
	CreateUser(userInfo model.User) error
}
type Repository struct {
	db *sql.DB
}

// NewRepository cria uma nova instância do Repository.
func NewRepository() (*Repository, error) {
	db, err := sql.Open("sqlite", "./db/main.db")
	if err != nil {
		return nil, fmt.Errorf("error opening sqlite database: %s", err)
	}
	return &Repository{db: db}, nil
}

func (r *Repository) CreateUser(userInfo model.User) error {
	stmt, err := r.db.Prepare("INSERT INTO user (first_name,last_name,user_name, password,email,) VALUES (?, ?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	// Execute SQL statement to insert data
	_, err = stmt.Exec(userInfo.FirstName, userInfo.LastName, userInfo.UserName, userInfo.Password, userInfo.Email)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *Repository) GetAll(r *UserRepository) []model.User {
	var users []User

	defer db.Close()

	allUsers, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer allUsers.Close()

	for allUsers.Next() {
		var user model.User
		if err := allUsers.Scan(&user.ID, &user.FirstName, &user.LastName, &model.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
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
