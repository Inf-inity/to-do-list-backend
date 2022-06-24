package server

import (
	"github.com/Inf-inity/to-do-list-backend/internal/pkg/backend/model"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (server *Server) addDB() error {
	db, err := gorm.Open(postgres.Open(server.config.GetString("db.dsn")), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "open database connection")
	}

	if err := db.AutoMigrate(&model.User{}, &model.Task{}, &model.Team{}); err != nil {
		return errors.Wrap(err, "auto migrate")
	}

	server.db = db

	return nil
}
