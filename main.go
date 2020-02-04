package main

import (
	"github.com/jub3i/go-webserver/routes"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting server on %s", ":1337")
	http.ListenAndServe(":1337", routes.Mux)
}
