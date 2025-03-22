package main

import (
	"log"
	"net/http"
	"server/auth"
	"server/config"
	"server/database"
	"server/handlers"
	"server/repositories"
	"server/services"
)

type App struct {
	config        *config.Config
	handlers      handlers.Handlers
	authenticator *auth.JWTAuthenticator
}

func (a *App) start() error {
	mux := http.NewServeMux()
	mux.Handle("POST /users/register", a.handlers.UserHandler.Register())
	mux.Handle("POST /users/login", a.handlers.UserHandler.Login())
	mux.Handle("GET /events", a.handlers.EventHandle.GetEvents())

	return http.ListenAndServe(a.config.ServerAddr, mux)
}

func main() {
	conf := config.NewConfig()
	db, err := database.Connect(&conf.DatabaseConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	authenticator := auth.NewJWTAuthenticator(conf.AuthConfig.JwtSecret, conf.AuthConfig.JwtIssuer)

	app := App{
		config:        conf,
		authenticator: authenticator,
		handlers: handlers.Handlers{
			UserHandler: handlers.NewDefaultUserHandler(
				services.NewDefaultUserService(
					repositories.NewDefaultUserRepository(db),
					authenticator,
				),
			),
			EventHandle: handlers.NewDefaultEventHandler(
				services.NewDefaultEventService(
					repositories.NewDefaultEventRepository(db),
				),
			),
		},
	}

	err = app.start()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
