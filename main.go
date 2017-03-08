package main

import (
	"fmt"
	"log"
	"net/http"

	"app/routes"
)

func main() {
	router := routes.NewRouter()

	fmt.Println("Server listen on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
