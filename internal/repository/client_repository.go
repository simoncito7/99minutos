package repository

const (
	_queryCreateUser = `INSERT INTO client (
		name,
		last_name,
		email,
		password,
		created_at,
		token
		)
		VALUES (?, ?, ?, ?, ?, ?);`
)
