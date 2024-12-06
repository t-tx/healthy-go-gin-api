package repositories

import (
	"database/sql"
	"fmt"
	"healthy/internal/database/models"
	"healthy/internal/pkg/utils"
	"log"
)

func CreateArticle(db *sql.DB, imageURLs, title, uploadedTime, tags, content string) (*models.Article, error) {
	query := `INSERT INTO articles (image_urls, title, uploaded_time, tags, content) 
              VALUES (?, ?, ?, ?, ?)`
	result, err := db.Exec(query, imageURLs, title, uploadedTime, tags, content)
	if err != nil {
		return nil, fmt.Errorf("failed to insert article: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	return &models.Article{
		ID:           int(lastInsertID),
		ImageURLs:    imageURLs,
		Title:        title,
		UploadedTime: uploadedTime,
		Content:      content,
		Tags:         tags,
		CreatedAt:    utils.GetTimeNowString(),
	}, nil
}

func ListArticles(db *sql.DB) ([]*models.Article, error) {
	rows, err := db.Query(`SELECT id, image_urls, title, uploaded_time, content, tags, created_at FROM articles`)
	if err != nil {
		return nil, fmt.Errorf("failed to query articles: %w", err)
	}
	defer rows.Close()

	var articles []*models.Article
	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.ImageURLs, &article.Title, &article.UploadedTime, &article.Tags, &article.Content, &article.CreatedAt); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		articles = append(articles, &article)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return articles, nil
}
