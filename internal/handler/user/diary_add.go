package user

import (
	"healthy/internal/database/models"
	"healthy/internal/database/repositories"
)

func (u *UserHandler) AddUserDiaries(req *AddUserDiariesReq) error {
	diaries := adaptUserDiaryToModel(req)
	err := repositories.AddUserDiaries(u.db, diaries)
	return err
}

type AddUserDiariesReq struct {
	Username string       `json:"username"`
	Diaries  []*UserDiary `json:"diaries"`
}

func adaptUserDiaryToModel(req *AddUserDiariesReq) []*models.UserDiary {
	var output = make([]*models.UserDiary, len(req.Diaries))
	for idx, diary := range req.Diaries {
		output[idx] = &models.UserDiary{
			Username: req.Username,
			Title:    diary.Title,
			Content:  diary.Content,
		}
	}
	return output
}
