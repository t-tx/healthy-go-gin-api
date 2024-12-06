package cmd

import (
	"healthy/internal/config"
	"healthy/internal/database"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

var (
	migrationConfigPath string
)

func init() {
	migrationCmd.PersistentFlags().StringVarP(&migrationConfigPath, "config", "c", "config.yml", "config to file path")
}

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "Migrate database",
	Long:  `Migrate database`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load(migrationConfigPath)

		if err != nil {
			log.Fatal().Err(err).Msg("Error message")
		}

		database.RunMigrations(cfg.Database.Path)
	},
}
