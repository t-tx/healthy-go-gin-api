package repositories

import (
	"database/sql"
	"fmt"
	"healthy/internal/database/models"
	"healthy/internal/pkg/utils"
	"strings"
)

const (
	batchSize int = 10
)

func AddUserEvents(db *sql.DB, events []*models.UserEvent) error {
	if len(events) == 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}
	var createdAt = utils.GetTimeNowString()
	baseQuery := `INSERT INTO user_events (username, event_type, content, created_at) VALUES `
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
			args = append(args, events[j].Username, events[j].EventType, events[j].Content, createdAt)
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

func ListUserEvent(db *sql.DB, username string) ([]*models.UserEvent, error) {
	query := `SELECT id, username, event_type, content, created_at FROM user_events 
              WHERE username = ? ORDER BY created_at DESC`

	rows, err := db.Query(query, username)
	if err != nil {
		return nil, fmt.Errorf("could not fetch events: %v", err)
	}
	defer rows.Close()

	var events []*models.UserEvent
	for rows.Next() {
		var event models.UserEvent
		var createdAt string
		err := rows.Scan(&event.ID, &event.Username, &event.EventType, &event.Content, &createdAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan event row: %v", err)
		}
		event.CreatedAt, _ = utils.ParseTime(createdAt)
		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("could not iterate over rows: %v", err)
	}

	return events, nil
}
