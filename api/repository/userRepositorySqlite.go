package repository

import (
	"database/sql"
	"fmt"

	"github.com/gateway-address/user"
	_ "modernc.org/sqlite"
)

type RepositorySqlite struct {
	db *sql.DB
}

func repositoryConfig(db *sql.DB) *sql.DB {
	db.Exec("PRAGMA foreign_keys = ON; PRAGMA timezone = 'UTC")
	return db
}

// NewRepository cria uma nova instância do Repository.
func NewRepositorySqlite() (*RepositorySqlite, error) {
	db, err := sql.Open("sqlite", "./db/main.db")
	if err != nil {
		return nil, fmt.Errorf("error opening sqlite database: %s", err)
	}
	db = repositoryConfig(db)
	return &RepositorySqlite{db: db}, nil
}

func (r *RepositorySqlite) Create(user *user.User) error {
	stmt, err := r.db.Prepare("INSERT INTO user (first_name, last_name, username, password, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute SQL statement to insert data
	_, error := stmt.Exec(user.FirstName, user.LastName, user.UserName, user.Password, user.Email)
	if error != nil {
		return error
	}
	return nil
}

func (r *RepositorySqlite) GetAll() ([]user.User, error) {
	var users []user.User

	userRows, err := r.db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println(err)
		return users, err // Return empty slice and error
	}
	defer userRows.Close()

	for userRows.Next() {
		var u user.User
		if err := userRows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	if err := userRows.Err(); err != nil {
		fmt.Println(err)
		return users, err // Return slice with retrieved users and error
	}
	return users, nil // Return slice of users and no error
}

func (r *RepositorySqlite) GetUserByID(id int) (user.User, error) {
	stmt, err := r.db.Prepare("SELECT * FROM user WHERE id = ?")
	if err != nil {
		return user.User{}, err
	}
	defer stmt.Close()

	var u user.User
	rows, err := stmt.Query(id)
	if err != nil {
		return u, err
	}
	defer rows.Close()

	if !rows.Next() {
		return u, fmt.Errorf("user with ID %d not found", id)
	}

	err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *RepositorySqlite) GetPaginatedUsers(limit int, offset int) ([]user.User, error) {
	var users []user.User
	userRows, err := r.db.Query("SELECT * FROM user LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		fmt.Println(err)
		return users, err // Return empty slice and error
	}
	defer userRows.Close()

	for userRows.Next() {
		var u user.User
		if err := userRows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	if err := userRows.Err(); err != nil {
		fmt.Println(err)
		return users, err // Return slice with retrieved users and error
	}
	return users, nil // Return slice of users and no error
}
