package main

import (
	"fmt"
	"log"
	"net/http"

	"app/config"
	"app/routes"

	"github.com/gorilla/handlers"
)

func main() {

	cfg := config.NewConfig()
	cfg.Load("config.yml")

	router := routes.NewRouter()

	fmt.Println("Server listen on http://localhost", cfg.Get("port"))
	log.Fatal(http.ListenAndServe(cfg.Get("port"), handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(router)))
}
