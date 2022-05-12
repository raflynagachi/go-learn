package app

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func (s *Server) InitializeCommands(arg string, appConfig AppConfig, dbConfig DBConfig) {
	if arg != "" {
		cmdApp := cli.NewApp()
		cmdApp.Commands = []*cli.Command{
			{
				Name: "db:migrate",
				Action: func(ctx *cli.Context) error {
					s.MigrateDB(false)
					return nil
				},
			},
			{
				Name: "db:migrate-fresh",
				Action: func(ctx *cli.Context) error {
					s.MigrateDB(true)
					return nil
				},
			},
			{
				Name: "db:seed",
				Action: func(ctx *cli.Context) error {
					s.SeedDB()
					return nil
				},
			},
			{
				Name: "start",
				Action: func(ctx *cli.Context) error {
					s.Run(":" + appConfig.AppPort)
					return nil
				},
			},
		}
		err := cmdApp.Run(os.Args)
		if err != nil {
			log.Fatal(err)
		}
	}
}
