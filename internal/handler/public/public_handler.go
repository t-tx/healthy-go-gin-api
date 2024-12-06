package public

import (
	"database/sql"
	"healthy/internal/database"
)

func New() *PublicHandler {
	return &PublicHandler{db: database.Get()}
}

type PublicHandler struct {
	db *sql.DB
}
