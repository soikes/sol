package crdb

const (
	createUsersStmt = `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			email STRING NOT NULL,
			name STRING NULL
		);`
	
	insertUserStmt = `
		INSERT INTO users (email, name) VALUES (
			$1, $2
		);`
)