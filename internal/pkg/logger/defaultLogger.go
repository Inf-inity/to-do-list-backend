package logger

import (
	"fmt"

	"github.com/kataras/golog"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// NewDefaultLogger returns a new default logger instance.
func NewDefaultLogger(displayName string, conf *viper.Viper) (*golog.Logger, error) {
	logger := golog.New()
	if err := ApplyDefaultOptions(logger, displayName, conf); err != nil {
		return nil, errors.Wrap(err, "apply default options")
	}

	return logger, nil
}

// ApplyDefaultOptions applies the default options to the logger instance.
func ApplyDefaultOptions(logger *golog.Logger, displayName string, conf *viper.Viper) error {
	if displayName != "" {
		logger.SetPrefix(fmt.Sprintf("[%s] ", displayName))
	}
	logger.SetTimeFormat(conf.GetString("timeFormat"))
	logger.SetLevel(conf.GetString("level"))
	fileWriter, err := NewLogFileWriter(conf)
	if err != nil {
		return errors.Wrap(err, "create log file writer")
	}
	logger.AddOutput(fileWriter)

	return nil
}
