package server

import (
	"encoding/json"
	"net/http"

	"github.com/Inf-inity/to-do-list-backend/internal/pkg/backend/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (server *Server) userByIDHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	u := &model.User{}
	server.db.First(&u, v["id"])
	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) newUserHandler(w http.ResponseWriter, r *http.Request) {
	inp := &model.UserInput{}
	if r.Body == http.NoBody {
		http.Error(w, "no body", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(inp); err != nil {
		http.Error(w, errors.Wrap(err, "unmarshal body").Error(), http.StatusBadRequest)
		return
	}
	u := &model.User{
		Name:  inp.Name,
		Teams: idArrayToTeamArray(inp.Teams),
	}
	server.db.Create(u)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func idArrayToTeamArray(ids []uint) []*model.Team {
	var teams []*model.Team
	for _, id := range ids {
		teams = append(teams, &model.Team{
			Model: gorm.Model{
				ID: id,
			},
		})
	}

	return teams
}

func idArrayToTaskArray(ids []uint) []*model.Task {
	var tasks []*model.Task
	for _, id := range ids {
		tasks = append(tasks, &model.Task{
			Model: gorm.Model{
				ID: id,
			},
		})
	}

	return tasks
}
