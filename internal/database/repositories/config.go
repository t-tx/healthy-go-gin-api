package repositories

import (
	"database/sql"
	"fmt"
	"healthy/internal/database/models"
	"log"
)

func CreateGlobalConfig(db *sql.DB, scope, key, value string) (*models.GlobalConfig, error) {
	query := `INSERT INTO global_config (scope, key, value) VALUES (?, ?, ?)`
	result, err := db.Exec(query, key, value)
	if err != nil {
		return nil, fmt.Errorf("failed to insert global config: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return &models.GlobalConfig{
		ID:    int(lastInsertID),
		Key:   key,
		Value: value,
	}, nil
}

func ListGlobalConfigs(db *sql.DB) ([]*models.GlobalConfig, error) {
	rows, err := db.Query(`SELECT id, key, value FROM global_config where scope = 'global'`)
	if err != nil {
		return nil, fmt.Errorf("failed to query global config: %w", err)
	}
	defer rows.Close()

	var configs []*models.GlobalConfig
	for rows.Next() {
		var config models.GlobalConfig
		if err := rows.Scan(&config.ID, &config.Key, &config.Value); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		configs = append(configs, &config)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return configs, nil
}
