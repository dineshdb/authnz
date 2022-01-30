package api

import (
	"database/sql"
	"net/http"

	"strconv"

	"github.com/dineshdb/authnz/internal/auth"
	"github.com/dineshdb/authnz/internal/middlewares"
	"github.com/dineshdb/authnz/internal/user"
	"github.com/dineshdb/authnz/internal/utils"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

/// App is the struct for holding the global configuration and state for the
/// application
type App struct {
	Config       AppConfig
	DB           sql.DB
	ArgonParams  utils.ArgonParams
	JWTValidator auth.JWTValidator
}

func (app *App) HandleRequests() {

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.RequestLoggerMiddleware(router))

	healthRouter := router.PathPrefix("/health").Subrouter().StrictSlash(true)
	healthRouter.HandleFunc("/live", Live)
	healthRouter.HandleFunc("/ready", Ready)

	apiRouter := router.PathPrefix("/api/v1").Subrouter().StrictSlash(true)

	privateRouter := apiRouter.PathPrefix("/").Subrouter().StrictSlash(true)
	// Register authentication middleware
	privateRouter.Use(app.JWTValidator.AuthMiddleware)
	privateRouter.HandleFunc("/profile/me", app.GetMyProfile)

	publicRouter := apiRouter.PathPrefix("/public").Subrouter().StrictSlash(true)
	publicRouter.HandleFunc("/auth/login", app.Login).Methods("POST")
	publicRouter.HandleFunc("/auth/register", app.Signup).Methods("POST")
	publicRouter.HandleFunc("/auth/refresh", app.Refresh).Methods("POST")

	// Database Migrations
	var err error = user.NewUserRepository(&app.DB).Migrate()
	if err != nil {
		log.Error().Msg("Couldn't migrate database table")
		return
	}

	var addr string = app.Config.Host + ":" + strconv.Itoa(app.Config.Port)
	log.Info().Msg("Listening on " + addr)
	log.Fatal().AnErr("error", (http.ListenAndServe(addr, router)))
}
