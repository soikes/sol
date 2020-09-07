package crdb

const (
	createUsersTableStmt = `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			email STRING NOT NULL UNIQUE,
			name STRING NOT NULL,
			password string NOT NULL 
		);`

	insertUserStmt = `
		INSERT INTO users (email, name, password) VALUES (
			$1, $2, $3
		) RETURNING id;`

	getUserPasswordStmt = `SELECT password FROM users WHERE email = $1;`

	getUserInfoStmt = `SELECT id, name, email FROM users where id = $1;`
)