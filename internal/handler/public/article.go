package public

import (
	"encoding/json"
	"healthy/internal/database/models"
	"healthy/internal/database/repositories"
	"strings"
)

func (p *PublicHandler) GetArticles() ([]*Article, error) {
	articles, err := repositories.ListArticles(p.db)
	if err != nil {
		return nil, err
	}
	return adaptArticles(articles)
}

type Article struct {
	ImageURLs    *ImgUrl  `json:"image_urls"`
	Title        string   `json:"title"`
	UploadedTime string   `json:"uploaded_time"`
	Tags         []string `json:"tags"`
}

type ImgUrl struct {
	BigImgUrl   string `json:"big_img_url"`
	SmallImgUrl string `json:"small_img_url"`
}

func adaptArticles(input []*models.Article) ([]*Article, error) {
	var output = make([]*Article, len(input))
	for idx, article := range input {
		var urls ImgUrl
		err := json.Unmarshal([]byte(article.ImageURLs), &urls)
		if err != nil {
			return nil, err
		}

		output[idx] = &Article{
			ImageURLs:    &urls,
			Title:        article.Title,
			UploadedTime: article.UploadedTime,
			Tags:         strings.Split(article.Tags, ","),
		}
	}

	return output, nil
}
