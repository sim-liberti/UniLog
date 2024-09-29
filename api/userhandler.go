package api

import (
	"net/http"
	"strconv"
	"teststack/utils"
)

func (s *Server) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users := s.database.GetUsers("")
	utils.WriteJSON(w, r, users)
}

func (s *Server) HandleGetUserByID(w http.ResponseWriter, r *http.Request) {
	requestID := r.PathValue("id")
	id, _ := strconv.Atoi(requestID)
	user := s.database.GetUserByID(id)
	utils.WriteJSON(w, r, user)
}

func (s *Server) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, r, nil)
}

func (s *Server) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, r, nil)
}

func (s *Server) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, r, nil)
}
