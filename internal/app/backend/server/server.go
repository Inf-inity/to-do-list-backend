package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Server struct {
	config *viper.Viper
	router *mux.Router
	db     *gorm.DB
}

func NewServer(conf *viper.Viper) (*Server, error) {
	server := &Server{
		config: conf,
		router: mux.NewRouter(),
	}

	if err := server.addDB(); err != nil {
		return nil, errors.Wrap(err, "add database connection")
	}

	server.router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("Hello world!"))
	})

	return server, nil
}

func (server *Server) ListenAndServe() error {
	return http.ListenAndServe(server.config.GetString("address"), server.router)
}
