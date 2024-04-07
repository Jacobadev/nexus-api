package repository

const (
	findUserByEmail = `SELECT *	FROM users WHERE email = $1`
)
