package main

import (
	"fmt"
	"net/http"
	"waysbuck/database"
	"waysbuck/pkg/mysql"
	"waysbuck/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    }

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// Initialization "uploads" folder to public here ...
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT","HEAD" ,"OPTIONS", "PATCH", "DETELE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = "5000"

	fmt.Println("server running localhost:" +port)
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}