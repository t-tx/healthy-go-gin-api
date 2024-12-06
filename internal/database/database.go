package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"healthy/internal/pkg/utils"
	"os"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const (
	scriptPath = "./scripts"
)

func Get() *sql.DB {
	return db
}

func Init(dbPath string) {
	var err error
	newDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to SQLite")
	}
	db = newDB
}

func RunMigrations(dbPath string) {
	files, err := os.ReadDir(scriptPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read directory")
	}

	var sqlScipts []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			sqlScipts = append(sqlScipts, file.Name())
		}
	}
	sort.Slice(sqlScipts, func(i, j int) bool {
		return sqlScipts[i] < sqlScipts[j]
	})

	newDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to SQLite")
	}
	db = newDB

	lastScript, err := getLastMigrateScriptFile()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get last migrate script file")
	}
	for _, filename := range sqlScipts {
		if lastScript < filename {
			scriptContent, err := os.ReadFile(fmt.Sprintf("%s/%s", scriptPath, filename))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to read file")
			}
			tx, err := db.Begin()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to begin transaction")
			}
			_, err = tx.Exec(string(scriptContent))
			if err != nil {
				log.Fatal().Err(err).Msg("failed to execute script")
			}
			var mv = &migrateValue{
				Content: string(scriptContent),
				File:    filename,
			}
			jsonValue, err := json.Marshal(mv)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to marshal json")
			}

			sqlStatement := `
				INSERT OR REPLACE INTO global_config (key, value, created_at, scope)
				VALUES
					('migrate', ?, ?, 'system');
			`

			_, err = tx.Exec(sqlStatement, jsonValue, utils.GetTimeNowString())
			if err != nil {
				log.Fatal().Err(err).Str("filename", filename).Msg("failed to insert global config")
			}
			err = tx.Commit()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to commit transaction")
			}

			log.Info().Str("file", filename).Msg("Executed")
		}
	}
}

type migrateValue struct {
	File    string `json:"file"`
	Content string `json:"content"`
}

func getLastMigrateScriptFile() (string, error) {
	var value string
	err := db.QueryRow(`SELECT value FROM global_config WHERE key = ? and scope = ? limit 1`, "migrate", "system").Scan(&value)
	if err != nil {
		if err.Error() == "sql: no rows in result set" || err.Error() == "no such table: global_config" {
			return "", nil
		}
		return "", err
	}

	var mv migrateValue
	err = json.Unmarshal([]byte(value), &mv)
	if err != nil {
		return "", err
	}
	return mv.File, nil
}
