package crdb

var (
	createUsersStmt = `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			email STRING NOT NULL,
			name STRING NULL
		)`
)