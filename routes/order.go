package routes

import (
	"waysbucks_BE/handlers"
	"waysbucks_BE/pkg/middleware"
	"waysbucks_BE/pkg/mysql"
	"waysbucks_BE/repositories"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	orderRepository := repositories.RepositoryOrder(mysql.DB)
	h := handlers.HandlerOrder(orderRepository)

	r.HandleFunc("/orders", middleware.Auth(h.FindOrders)).Methods("GET")
	r.HandleFunc("/order/{id}", middleware.Auth(h.GetOrder)).Methods("GET")
	// Create "/order" route using middleware Auth, middleware UploadFile, handler CreateOrder, and method POST
	r.HandleFunc("/order/{id}", middleware.Auth(h.CreateOrder)).Methods("POST")
	r.HandleFunc("/order/{id}", middleware.Auth(h.DeleteOrder)).Methods("DELETE")
	// r.HandleFunc("/order/{id}", middleware.Auth(h.UpdateOrder)).Methods("PATCH")
}
