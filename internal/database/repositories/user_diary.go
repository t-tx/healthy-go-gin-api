package repositories

import (
	"database/sql"
	"fmt"
	"healthy/internal/database/models"
	"healthy/internal/pkg/utils"
	"strings"
)

func AddUserDiaries(db *sql.DB, events []*models.UserDiary) error {
	if len(events) == 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	var createdAt = utils.GetTimeNowString()

	baseQuery := `INSERT INTO user_diaries (username, title, content, created_at) VALUES `
	valuePlaceholders := `(?, ?, ?, ?)`

	for i := 0; i < len(events); i += batchSize {
		end := i + batchSize
		if end > len(events) {
			end = len(events)
		}

		var queryBuilder strings.Builder
		queryBuilder.WriteString(baseQuery)

		args := make([]interface{}, 0, (end-i)*3)
		for j := i; j < end; j++ {
			if j > i {
				queryBuilder.WriteString(", ")
			}
			queryBuilder.WriteString(valuePlaceholders)
			args = append(args, events[j].Username, events[j].Title, events[j].Content, createdAt)
		}

		query := queryBuilder.String()
		if _, err := tx.Exec(query, args...); err != nil {
			tx.Rollback()
			return fmt.Errorf("could not insert events: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}
func ListUserDiary(db *sql.DB, username string) ([]*models.UserDiary, error) {
	query := `SELECT id, username, title, content, created_at FROM user_diaries 
              WHERE username = ? ORDER BY created_at DESC`

	rows, err := db.Query(query, username)
	if err != nil {
		return nil, fmt.Errorf("could not fetch diaries: %v", err)
	}
	defer rows.Close()

	var diaries []*models.UserDiary
	for rows.Next() {
		var diary models.UserDiary
		var created_at string
		err := rows.Scan(&diary.ID, &diary.Username, &diary.Title, &diary.Content, &created_at)
		if err != nil {
			return nil, fmt.Errorf("could not scan diary row: %v", err)
		}
		diary.CreatedAt, _ = utils.ParseTime(created_at)
		diaries = append(diaries, &diary)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over rows: %v", err)
	}

	return diaries, nil
}
