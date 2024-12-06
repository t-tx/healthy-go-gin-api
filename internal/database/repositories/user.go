package repositories

import (
	"database/sql"
	"healthy/internal/database/models"
	"healthy/internal/pkg/defined"
	"healthy/internal/pkg/utils"
)

func AddUser(db *sql.DB, user *models.User) error {
	query := "INSERT INTO users (username, gender, birthday, password, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, user.Username, user.Gender, user.Birthday, user.Password, utils.GetTimeNowString())
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			return defined.ErrUsernameExists
		}
		return err
	}

	return nil
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	query := "SELECT id, username, gender, birthday, password FROM users WHERE username = ?"
	row := db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Gender, &user.Birthday, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
