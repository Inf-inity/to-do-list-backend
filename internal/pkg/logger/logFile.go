package logger

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

// LogFileWriter handles writing log files.
type LogFileWriter struct {
	mu          sync.Mutex
	currentFile *os.File
	config      *viper.Viper
	cron        *cron.Cron
}

// NewLogFileWriter returns a new LogFileWriter instance.
func NewLogFileWriter(conf *viper.Viper) (*LogFileWriter, error) {
	writer := &LogFileWriter{
		config: conf,
		cron:   cron.New(),
	}

	if err := writer.refreshCurrentFile(); err != nil {
		return nil, errors.Wrap(err, "refresh log file")
	}
	if err := writer.scheduleRefreshes(); err != nil {
		return nil, errors.Wrap(err, "schedule log file refreshes")
	}
	writer.cron.Start()

	return writer, nil
}

// scheduleRefreshes schedules the refresh of the log file to use.
func (writer *LogFileWriter) scheduleRefreshes() error {
	_, err := writer.cron.AddFunc(writer.config.GetString("refreshSpec"), func() {
		_ = writer.refreshCurrentFile()
	})
	return err
}

// refreshCurrentFile refreshes the used log file.
func (writer *LogFileWriter) refreshCurrentFile() error {
	writer.mu.Lock()
	defer writer.mu.Unlock()

	logDirPath := writer.config.GetString("logDir")
	if err := os.MkdirAll(logDirPath, os.FileMode(writer.config.GetInt("logDirMode"))); err != nil {
		return errors.Wrap(err, "create log dir")
	}

	fileName := time.Now().Format(writer.config.GetString("logFileNameFormat")) + ".log"
	f, err := os.OpenFile(
		filepath.Join(logDirPath, fileName),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		os.FileMode(writer.config.GetInt("logFileMode")),
	)
	if err != nil {
		return errors.Wrap(err, "create log file")
	}
	_ = writer.currentFile.Close()
	writer.currentFile = f

	return nil
}

// Write writes the passed date to the log file.
func (writer *LogFileWriter) Write(b []byte) (n int, err error) {
	writer.mu.Lock()
	defer writer.mu.Unlock()
	return writer.currentFile.Write(b)
}
