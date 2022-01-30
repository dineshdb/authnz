package main

import (
	"database/sql"

	"github.com/dineshdb/authnz/internal/api"
	"github.com/dineshdb/authnz/internal/auth"
	"github.com/dineshdb/authnz/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// TODO: Add configuration management. Use [12-factor application](12factor.net) guidelines for better maintainability of code

	// Using structured logs for easier log analysis. The output is [ndjson](ndjson.org)
	zerolog.SetGlobalLevel(zerolog.TraceLevel)

	var config api.AppConfig = api.AppConfig{
		Host:        "0.0.0.0",
		Port:        8080,
		DatabaseUrl: "./db.sqlite",
	}

	db, err := sql.Open("sqlite3", config.DatabaseUrl)
	if err != nil {
		log.Error().Msg("Got error opening database")
		return
	}

	var app api.App = api.App{
		Config:       config,
		DB:           *db,
		ArgonParams:  utils.DefaultArgonParams,
		JWTValidator: auth.New("private.pem"),
	}

	app.HandleRequests()
}
