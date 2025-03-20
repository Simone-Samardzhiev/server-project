package main

import (
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/handlers"
	"server/repositories"
	"server/services"
)

type App struct {
	config   *config.Config
	handlers handlers.Handlers
}

func (a *App) start() error {
	mux := http.NewServeMux()
	mux.Handle("POST /users/register", a.handlers.UserHandler.Register())

	return http.ListenAndServe(a.config.ServerAddr, mux)
}

func main() {
	conf := config.NewConfig()
	db, err := database.Connect(&conf.DatabaseConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	app := App{
		config: conf,
		handlers: handlers.Handlers{
			UserHandler: handlers.NewDefaultUserHandler(
				services.NewDefaultUserService(
					repositories.NewDefaultUserRepository(db),
				),
			),
		},
	}

	err = app.start()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
