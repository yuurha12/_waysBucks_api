package routes

import (
	"ways-bucks-api/handlers"
	"ways-bucks-api/pkg/mysql"
	"ways-bucks-api/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
