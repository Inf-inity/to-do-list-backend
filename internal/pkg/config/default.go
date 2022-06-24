package config

import (
	"path/filepath"

	"github.com/Inf-inity/to-do-list-backend/internal/pkg/config/defaults"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// LoadDefaultConfig loads a default configuration.
func LoadDefaultConfig(appName string, defaultValues ...map[string]interface{}) (*viper.Viper, error) {
	instance := NewDefaultConfig(appName)
	defaults.Set(instance, defaultValues...)
	if err := instance.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "read config from file")
	}

	return instance, nil
}

// NewDefaultConfig creates a new default configuration instance.
func NewDefaultConfig(appName string) *viper.Viper {
	instance := viper.New()
	ApplyDefaultOptions(instance, appName)
	return instance
}

// ApplyDefaultOptions applies the default values to the configuration instance.
func ApplyDefaultOptions(instance *viper.Viper, appName string) {
	ApplyDefaultSearchOptions(instance, appName)
	ApplyDefaultEnvOptions(instance, appName)
	ApplyDefaultWatchOptions(instance)
}

// ApplyDefaultWatchOptions applies the default watch options to the configuration instance.
func ApplyDefaultWatchOptions(instance *viper.Viper) {
	instance.WatchConfig()
}

// ApplyDefaultEnvOptions applies the default environment options to the configuration instance.
func ApplyDefaultEnvOptions(instance *viper.Viper, appName string) {
	instance.SetEnvPrefix(strcase.ToScreamingSnake("Portal " + appName))
	instance.AutomaticEnv()
}

// ApplyDefaultSearchOptions applies the default search options to the configuration instance.
func ApplyDefaultSearchOptions(instance *viper.Viper, appName string) {
	appName = strcase.ToKebab(appName)
	instance.SetConfigName(appName + ".conf")
	instance.AddConfigPath(filepath.Join("/etc/inventory", appName))
	instance.AddConfigPath(filepath.Join("$HOME/.inventory", appName))
	instance.AddConfigPath(".")
}
