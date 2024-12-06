package handler

import (
	"healthy/internal/handler/public"
	"healthy/internal/handler/user"
	"healthy/internal/pkg/enc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthy struct {
	user   *user.UserHandler
	public *public.PublicHandler
}

func (h *Healthy) AddUserDiaries(c *gin.Context) {
	var diaries []*user.UserDiary
	if err := c.ShouldBindJSON(&diaries); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	var req = &user.AddUserDiariesReq{
		Username: username.(string),
		Diaries:  diaries,
	}
	err := h.user.AddUserDiaries(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h *Healthy) GetUserDiaries(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	diaries, err := h.user.GetUserDiaries(&user.GetUserDiariesReq{Username: username.(string)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    diaries,
	})
}

func (h *Healthy) AddUserEvents(c *gin.Context) {

	var events []*user.UserEvent
	if err := c.ShouldBindJSON(&events); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	err := h.user.AddUserEvents(&user.AddUserEventReq{
		UserEvents: events,
		Username:   username.(string),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h *Healthy) GetUserEvents(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}
	userEvents, err := h.user.GetUserEvents(&user.GetUserEventsReq{
		Username: username.(string),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userEvents,
	})
}

func (h *Healthy) GetReportAchievements(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	achievement, err := h.user.GetReportAchievement(&user.GetReportAchievementsReq{Username: username.(string)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    achievement,
	})
}

func (h *Healthy) GetReportGraphs(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}

	graphs, err := h.user.GetReportGraphs(&user.GetReportGraphsReq{Username: username.(string)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    graphs,
	})
}

func (h *Healthy) GetArticles(c *gin.Context) {
	articles, err := h.public.GetArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    articles,
	})
}

func (h *Healthy) GetGlobalConfigs(c *gin.Context) {
	configs, err := h.public.GetGlobalConfigs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    configs,
	})
}

func (h *Healthy) SignUp(c *gin.Context) {
	var req user.SignUpReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	err := h.user.SignUp(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h *Healthy) SignIn(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	user, err := h.user.SignIn(&user.SignInReq{
		Username: credentials.Username,
		Password: credentials.Password,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := enc.GenerateJWT(credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.SetCookie("token", token, 24*60*60, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"success": true, "token": token, "user": user})
}
