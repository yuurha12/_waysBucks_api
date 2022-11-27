package routes

import (
	"waysbucks_BE/handlers"
	"waysbucks_BE/pkg/middleware"
	"waysbucks_BE/pkg/mysql"
	"waysbucks_BE/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	r.HandleFunc("/products", h.FindProducts).Methods("GET")
	r.HandleFunc("/product/{id}", middleware.Auth(h.GetProduct)).Methods("GET")
	// Create "/product" route using middleware Auth, middleware UploadFile, handler CreateProduct, and method POST
	r.HandleFunc("/product", middleware.Auth(middleware.UploadFile(h.CreateProduct))).Methods("POST")
	r.HandleFunc("/product/{id}", middleware.Auth(h.DeleteProduct)).Methods("DELETE")
	r.HandleFunc("/product/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProduct))).Methods("PATCH")
}
