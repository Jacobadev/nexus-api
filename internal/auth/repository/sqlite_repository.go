package repository

import (
	model "github.com/gateway-address/internal/models"
	"github.com/gateway-address/pkg/utils"
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
	if err := r.db.QueryRowx(getUserByEmailQuery, &user.Email).StructScan(foundUser); err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (r *RepositorySqlite) GetByID(user_id int) (*model.User, error) {
	foundUser := &model.User{}
	if err := r.db.QueryRowx(getUserByIDQuery, &user_id).StructScan(foundUser); err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (r *RepositorySqlite) Delete(id int) error {
	_, err := r.db.Exec(deleteUserByIDQuery, &id)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositorySqlite) GetUsers(pq *utils.PaginationQuery) (*model.UsersList, error) {
	var totalCount int
	if err := r.db.Get(&totalCount, getTotalQuery); err != nil {
		return nil, err
	}

	if totalCount == 0 {
		return &model.UsersList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			Users:      make([]*model.User, 0),
		}, nil
	}

	users := make([]*model.User, 0, pq.GetSize())
	if err := r.db.Select(
		&users,
		getUsersQuery,
		pq.GetOrderBy(),
		pq.GetOffset(),
		pq.GetLimit(),
	); err != nil {
		return nil, err
	}

	return &model.UsersList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		Users:      users,
	}, nil
}

//	func (r *RepositorySqlite) GetPaginated(limit int, offset int) ([]model.User, error) {
//		var users []model.User
//		userRows, err := r.db.Query("SELECT * FROM user LIMIT ? OFFSET ?", limit, offset)
//		if err != nil {
//			fmt.Println(err)
//			return users, err // Return empty slice and error
//		}
//		defer userRows.Close()
//
//		for userRows.Next() {
//			var u model.User
//			if err := userRows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
//				fmt.Println(err)
//				continue
//			}
//			users = append(users, u)
//		}
//		if err := userRows.Err(); err != nil {
//			fmt.Println(err)
//			return users, err // Return slice with retrieved users and error
//		}
//		return users, nil // Return slice of users and no error

//	}

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
