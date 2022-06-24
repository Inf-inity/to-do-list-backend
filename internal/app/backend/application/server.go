package application

import (
	"github.com/Inf-inity/to-do-list-backend/internal/app/backend/server"
	"github.com/pkg/errors"
)

func (app *App) addServer() {
	s, err := server.NewServer(app.Config.Sub("server"))
	if err != nil {
		app.Logger.Fatal(errors.Wrap(err, "new server"))
	}

	app.Server = s
}
