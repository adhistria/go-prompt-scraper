package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adhistria/go-prompt-scraper/metadata/repository/file"
	"github.com/adhistria/go-prompt-scraper/metadata/repository/rest"

	"github.com/adhistria/go-prompt-scraper/metadata/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/urfave/cli/v2"
)

func main() {
	mjr := file.NewMetadataJSONRepository()
	srp := rest.NewScraperRepository()
	s := service.NewService(srp, mjr)
	var site string
	app := &cli.App{
		Name:  "fetch",
		Usage: "get metadata of site that you want!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "metadata",
				Aliases:     []string{"m"},
				Usage:       "Load metadata of site that you had access",
				Destination: &site,
			},
		},
		Action: func(ctx *cli.Context) error {
			if site != "" {
				err := s.FetchDetail(ctx.Context, site)
				if err != nil {
					return err
				}
			} else {
				err := s.Fetch(ctx.Context, ctx.Args().Slice())
				if err != nil {
					return err

				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func initPostgres() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost, dbPort)

	db, err := sqlx.Connect(dbDriver, connection)
	if err != nil {
		return nil, err
	}
	return db, nil

}
