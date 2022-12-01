package routes

import (
	"ways-bucks-api/handlers"
	"ways-bucks-api/pkg/middleware"
	"ways-bucks-api/pkg/mysql"
	"ways-bucks-api/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authReposutory := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authReposutory)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
