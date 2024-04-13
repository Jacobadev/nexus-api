package repository

const (
	getUserByEmailQuery = `SELECT user_id, username,password,  first_name,last_name,email	FROM users WHERE email = $1`

	createUserQuery = `INSERT INTO users (username, password, email, first_name, last_name, created_at, updated_at)
					VALUES ($1, $2, $3, $4, $5, now(), now()) RETURNING *`
	getUserByIDQuery    = `SELECT user_id, username, password,first_name,last_name,email	FROM users WHERE user_id = $1`
	deleteUserByIDQuery = `DELETE FROM users WHERE user_id = $1`

	getTotalQuery = `SELECT COUNT(user_id) FROM users`

	getUsersQuery = `SELECT user_id, first_name, last_name, email, 
       			 created_at, updated_at 
				 FROM users 
				 ORDER BY COALESCE(NULLIF($1, ''), first_name) OFFSET $2 LIMIT $3`
)
