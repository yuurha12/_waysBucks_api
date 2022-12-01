package routes

import (
	"ways-bucks-api/handlers"
	"ways-bucks-api/pkg/middleware"
	"ways-bucks-api/pkg/mysql"
	"ways-bucks-api/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	r.HandleFunc("/profiles", h.FindProfiles).Methods("GET")
	r.HandleFunc("/profile/{id}", h.GetProfile).Methods("GET")
	r.HandleFunc("/profile", h.CreateProfile).Methods("POST")
	r.HandleFunc("/profile", middleware.Auth(middleware.UploadFile(h.UpdateProfile))).Methods("PATCH")
	r.HandleFunc("/profile/{id}", h.DeleteProfile).Methods("DELETE")
}
