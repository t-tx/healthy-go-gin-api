package handler

import (
	"healthy/internal/handler/public"
	"healthy/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func NewHealthyHandler() IHealthy {
	return &Healthy{
		user:   user.New(),
		public: public.New(),
	}
}

type IHealthy interface {
	AddUserDiaries(c *gin.Context)
	GetUserDiaries(c *gin.Context)
	AddUserEvents(c *gin.Context)
	GetUserEvents(c *gin.Context)
	GetReportAchievements(c *gin.Context)
	GetReportGraphs(c *gin.Context)
	GetArticles(c *gin.Context)
	GetGlobalConfigs(c *gin.Context)
	SignIn(c *gin.Context)
	SignUp(c *gin.Context)
}
