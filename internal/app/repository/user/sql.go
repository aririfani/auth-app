package user

const (
	QueryCreateUser = `
		INSERT INTO users(username, phone, registered_at, password, role) VALUES($1, $2, $3, $4, $5)
	`

	QueryGetUserByField = `SELECT * from users where LOWER(%s) = LOWER(%s) ORDER BY registered_at limit 1`
)
