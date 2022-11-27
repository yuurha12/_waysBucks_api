package routes

import (
	"waysbucks_BE/handlers"
	"waysbucks_BE/pkg/middleware"
	"waysbucks_BE/pkg/mysql"
	"waysbucks_BE/repositories"

	"github.com/gorilla/mux"
)

func ToppingRoutes(r *mux.Router) {
	toppingRepository := repositories.RepositoryTopping(mysql.DB)
	h := handlers.HandlerTopping(toppingRepository)

	r.HandleFunc("/toppings", h.FindToppings).Methods("GET")
	r.HandleFunc("/topping/{id}",middleware.Auth (h.GetTopping)).Methods("GET")
	// Create "/topping" route using middleware Auth, middleware UploadFile, handler Createtopping, and method POST
	r.HandleFunc("/topping", middleware.Auth(middleware.UploadFile(h.CreateTopping))).Methods("POST")
	r.HandleFunc("/topping/{id}", h.DeleteTopping).Methods("DELETE")
	r.HandleFunc("/topping/{id}", middleware.UploadFile(h.UpdateTopping)).Methods("PATCH")
}
