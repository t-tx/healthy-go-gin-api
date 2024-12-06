package cmd

import (
	"fmt"
	v1 "healthy/internal/api/v1"
	"healthy/internal/config"
	"healthy/internal/database"
	"healthy/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

var (
	configPath string
)

func init() {
	serveCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config.yml", "config to file path")
}

var serveCmd = &cobra.Command{
	Use:   "healthy",
	Short: "Start healthy application",
	Long:  `Start healthy application`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load(configPath)

		if err != nil {
			log.Fatal().Msg("Error message")
		}
		database.Init(cfg.Database.Path)
		lvl, _ := zerolog.ParseLevel(cfg.Log.Level)
		zerolog.SetGlobalLevel(lvl)

		r := gin.Default()
		healthyHandler := handler.NewHealthyHandler()

		v1.RegisterRoutes(r, healthyHandler)

		r.Run(fmt.Sprintf(":%s", cfg.HTTP.Port))
	},
}
