package user

import (
	"database/sql"
	"healthy/internal/database"
)

func New() *UserHandler {
	return &UserHandler{
		db: database.Get(),
	}
}

type UserHandler struct {
	db *sql.DB
}
