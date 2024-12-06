package user

import (
	"healthy/internal/database/repositories"
	"time"
)

func (u *UserHandler) GetUserDiaries(req *GetUserDiariesReq) ([]*UserDiary, error) {
	diaries, err := repositories.ListUserDiary(u.db, req.Username)
	if err != nil {
		return nil, err
	}
	var output = make([]*UserDiary, len(diaries))
	for idx, diary := range diaries {
		output[idx] = &UserDiary{
			Username:  diary.Username,
			CreatedAt: diary.CreatedAt.Format(time.RFC3339),
			Title:     diary.Title,
			Content:   diary.Content,
		}
	}
	return output, nil
}

type GetUserDiariesReq struct {
	Username string `json:"username"`
}
