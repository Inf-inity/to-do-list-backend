package application

import (
	"log"

	"github.com/Inf-inity/to-do-list-backend/internal/app/backend/server"
	"github.com/Inf-inity/to-do-list-backend/internal/pkg/application"
	"github.com/Inf-inity/to-do-list-backend/internal/pkg/config/defaults"
)

// App is used to bundle all application resources.
type App struct {
	application.DefaultApp

	Server *server.Server
}

// NewApp returns a new App instance.
func NewApp() *App {
	app := &App{}

	if err := app.Init(application.Manifest{
		Name:           "backend",
		DisplayName:    "ToDo-List Backend",
		ConfigDefaults: []map[string]interface{}{defaults.Application},
	}); err != nil {
		log.Panicln(err)
	}

	return app
}
