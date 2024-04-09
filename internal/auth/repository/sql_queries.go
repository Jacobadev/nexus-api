package repository

const (
	findUserByEmail = `SELECT user_id, username, password, first_name,last_name,email	FROM users WHERE email = $1`

	createUserQuery = `INSERT INTO users (username, password, email, first_name, last_name, created_at, updated_at)
					VALUES ($1, $2, $3, $4, $5, now(), now()) RETURNING *`
)
