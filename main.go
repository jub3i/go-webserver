package main

import (
	"github.com/jub3i/go-webserver/config"
	"github.com/jub3i/go-webserver/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// load config
	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// start the http server
	addr := "localhost:" + os.Getenv("GOWS_PORT")
	log.Printf("starting go-webserver on `%s`", addr)
	http.ListenAndServe(addr, routes.Mux)
}
