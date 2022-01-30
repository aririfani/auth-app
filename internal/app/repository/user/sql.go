package user

const (
	QueryCreateUser = `
		INSERT INTO users(username, phone, registered_at, password, role) VALUES($1, $2, $3, $4, $5)
	`

	QueryGetUserByUsername = `SELECT * from users where LOWER(username) = LOWER($1)`
)
