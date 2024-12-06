package v1

import (
	"healthy/internal/config"
	"healthy/internal/handler"
	"healthy/internal/pkg/enc"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, healthy handler.IHealthy) {
	v1 := router.Group("/api/v1")

	v1.GET("/config/global", healthy.GetGlobalConfigs)

	v1.POST("/signin", healthy.SignIn)
	v1.POST("/signup", healthy.SignUp)

	v1.GET("/articles/health", healthy.GetArticles)

	user := v1.Group("/user")
	user.Use(authMiddleware())

	user.POST("/events", healthy.AddUserEvents)
	user.GET("/events", healthy.GetUserEvents)
	user.POST("/diaries", healthy.AddUserDiaries)
	user.GET("/diaries", healthy.GetUserDiaries)
	user.GET("/graph", healthy.GetReportGraphs)
	user.GET("/achievement", healthy.GetReportAchievements)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, no token provided"})
			c.Abort()
			return
		}

		claims := &enc.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			var jwtSecret = []byte(config.Get().App.JwtSecretKey)
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
