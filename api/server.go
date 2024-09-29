package api

import (
	"net/http"
	"teststack/database"
)

type Server struct {
	listenAddr string
	database   database.Database
}

func NewServer(listenAddr string, db database.Database) *Server {
	return &Server{
		listenAddr: listenAddr,
		database:   db,
	}
}

func (s *Server) Start() error {
	router := http.NewServeMux()
	// USER CRUD OPERATIONS
	router.HandleFunc("GET /user", s.HandleGetUsers)
	router.HandleFunc("GET /user/{id}", s.HandleGetUserByID)
	router.HandleFunc("PUT /user", s.HandleCreateUser)
	router.HandleFunc("POST /user/{id}", s.HandleUpdateUser)
	router.HandleFunc("DELETE /user/{id}", s.HandleDeleteUser)

	server := http.Server{
		Addr:    s.listenAddr,
		Handler: router,
	}

	return server.ListenAndServe()
}
