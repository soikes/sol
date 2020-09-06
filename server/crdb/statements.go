package crdb

const (
	createUsersStmt = `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			email STRING NOT NULL UNIQUE,
			name STRING NOT NULL,
			password string NOT NULL 
		);`
	
	insertUserStmt = `
		INSERT INTO users (email, name, password) VALUES (
			$1, $2, $3
		);`

	getUserPasswordStmt = `SELECT password FROM users WHERE email = $1;`
)