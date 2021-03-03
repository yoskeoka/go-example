package repository

import "database/sql"

// UserRepo implements domain.User to process user's personal data.
type UserRepo struct {
	conn *sql.DB
}

// TODO: Implement domain.User interface.
