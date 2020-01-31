package main

import (
	"log"
	"net/http"
	routes "webserver/routes"
)

func main() {
	log.Printf("Starting server on %s", ":1337")
	http.ListenAndServe(":1337", routes.Mux)
}
