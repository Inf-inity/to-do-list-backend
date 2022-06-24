package server

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Server struct {
	router *mux.Router
}

func NewServer(conf *viper.Viper) (*Server, error) {
	return &Server{router: mux.NewRouter()}, nil
}
