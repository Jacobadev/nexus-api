package repository

import (
	model "github.com/gateway-address/internal/models"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type RepositorySqlite struct {
	db *sqlx.DB
}

func NewRepositorySqlite(db *sqlx.DB) *RepositorySqlite {
	return &RepositorySqlite{db: db}
}

// NewRepository cria uma nova instância do Repository.
func (r *RepositorySqlite) Register(user *model.User) (*model.User, error) {
	u := &model.User{}

	if err := r.db.QueryRowx(createUserQuery, &user.UserName, &user.Password, &user.Email, &user.FirstName, &user.LastName).StructScan(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (r *RepositorySqlite) FindByEmail(user *model.User) (*model.User, error) {
	foundUser := &model.User{}
	if err := r.db.QueryRowx(findUserByEmail, &user.Email).StructScan(foundUser); err != nil {
		return nil, err
	}
	return foundUser, nil
}

// func (r *RepositorySqlite) GetAll() ([]model.User, error) {
// 	var users []model.User
//
// 	userRows, err := r.db.Query("SELECT * FROM user")
// 	if err != nil {
// 		fmt.Println(err)
// 		return users, err // Return empty slice and error
// 	}
// 	defer userRows.Close()
//
// 	for userRows.Next() {
// 		var u model.User
// 		if err := userRows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		users = append(users, u)
// 	}
// 	if err := userRows.Err(); err != nil {
// 		fmt.Println(err)
// 		return users, err // Return slice with retrieved users and error
// 	}
// 	return users, nil // Return slice of users and no error
// }
//
// func (r *RepositorySqlite) GetByID(id int) (model.User, error) {
// 	stmt, err := r.db.Prepare("SELECT * FROM user WHERE id = ?")
// 	if err != nil {
// 		return model.User{}, err
// 	}
// 	defer stmt.Close()
//
// 	var u model.User
// 	rows, err := stmt.Query(id)
// 	if err != nil {
// 		return u, err
// 	}
// 	defer rows.Close()
//
// 	if !rows.Next() {
// 		return u, fmt.Errorf("user with ID %d not found", id)
// 	}
//
// 	err = rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
// 	if err != nil {
// 		return u, err
// 	}
//
// 	return u, nil
// }
//
// func (r *RepositorySqlite) GetPaginated(limit int, offset int) ([]model.User, error) {
// 	var users []model.User
// 	userRows, err := r.db.Query("SELECT * FROM user LIMIT ? OFFSET ?", limit, offset)
// 	if err != nil {
// 		fmt.Println(err)
// 		return users, err // Return empty slice and error
// 	}
// 	defer userRows.Close()
//
// 	for userRows.Next() {
// 		var u model.User
// 		if err := userRows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		users = append(users, u)
// 	}
// 	if err := userRows.Err(); err != nil {
// 		fmt.Println(err)
// 		return users, err // Return slice with retrieved users and error
// 	}
// 	return users, nil // Return slice of users and no error
// }
//
// func (r *RepositorySqlite) DeleteByID(id int) error {
// 	_, err := r.db.Exec("DELETE FROM user WHERE id = ?", id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (r *RepositorySqlite) UpdateByID(id int, user *model.User) error {
// 	_, err := r.db.Exec("UPDATE user SET first_name = ?, last_name = ?, username = ?, email = ? WHERE id = ?", user.FirstName, user.LastName, user.UserName, user.Email, user.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (r *RepositorySqlite) Close() error {
// 	return r.db.Close()
// }
//
// func (r *RepositorySqlite) PartialUpdateByID(id int, user *model.User) error {
// 	_, err := r.db.Exec("UPDATE user SET first_name = ?, last_name = ?, username = ?, email = ? WHERE id = ?", user.FirstName, user.LastName, user.UserName, user.Email, user.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (r *RepositorySqlite) Stats() sql.DBStats {
// 	stats := r.db.Stats()
// 	return stats
