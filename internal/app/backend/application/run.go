package application

import "github.com/pkg/errors"

func (app *App) Run() {
	if err := app.Server.ListenAndServe(); err != nil {
		app.Logger.Fatal(errors.Wrap(err, "listen and serve"))
	}
}
