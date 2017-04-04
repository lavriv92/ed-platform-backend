package main

import (
	"fmt"
	"log"
	"net/http"

	"app/routes"
	"github.com/gorilla/handlers"
)

func main() {
	router := routes.NewRouter()

	fmt.Println("Server listen on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST"}),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)(router)))
}
