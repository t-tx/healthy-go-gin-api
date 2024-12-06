package user

import (
	"errors"
	"healthy/internal/pkg/defined"
	"healthy/internal/pkg/utils"
	"time"
)

func (u *UserHandler) GetReportAchievement(req *GetReportAchievementsReq) (*Achievement, error) {
	events, err := u.GetUserEvents(&GetUserEventsReq{Username: req.Username})
	if err != nil {
		return nil, errors.New("could not get user events")
	}
	exerciseCounter := 0
	lastMonth := time.Now().AddDate(0, -1, 0)
	lastMonthStr := utils.GetTimeString(lastMonth)
	for _, event := range events {
		if lastMonthStr > event.CreatedAt {
			break
		}

		if event.EventType == defined.EXERCISE {
			exerciseCounter++
		}
	}
	return &Achievement{
		Value: int(exerciseCounter * 100 / 30),
		Time:  utils.GetTimeNowString(),
	}, nil
}

type Achievement struct {
	Value int    `json:"value"`
	Time  string `json:"time"`
}
type GetReportAchievementsReq struct {
	Username string `json:"username"`
}
