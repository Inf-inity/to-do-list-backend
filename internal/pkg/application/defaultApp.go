package application

import (
	"github.com/Inf-inity/to-do-list-backend/internal/pkg/config"
	"github.com/Inf-inity/to-do-list-backend/internal/pkg/logger"
	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// DefaultApp bundles the default app resources.
type DefaultApp struct {
	Manifest Manifest
	Config   *viper.Viper
	Logger   *golog.Logger
}

// Init initializes the DefaultApp.
func (app *DefaultApp) Init(manifest Manifest) error {
	app.Manifest = manifest

	if err := app.addConfig(); err != nil {
		return errors.Wrap(err, "add configuration")
	}
	if err := app.addLogger(); err != nil {
		return errors.Wrap(err, "add logger")
	}

	return nil
}

// addConfig initializes the configuration to use by the application during app initialization.
func (app *DefaultApp) addConfig() error {
	c, err := config.LoadDefaultConfig(app.Manifest.Name, app.Manifest.ConfigDefaults...)
	if err != nil {
		return err
	}
	app.Config = c

	return nil
}

func (app *DefaultApp) addLogger() error {
	l, err := logger.NewDefaultLogger(app.Manifest.DisplayName, app.Config.Sub("logger"))
	if err != nil {
		return err
	}
	app.Logger = l

	return nil
}
