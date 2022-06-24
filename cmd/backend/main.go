package main

import "github.com/Inf-inity/to-do-list-backend/internal/app/backend/application"

func main() {
	app := application.NewApp()

	app.Logger.Info("hello world!")
}
