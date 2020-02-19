package main

import (
	"github.com/jub3i/go-webserver/config"
	"github.com/jub3i/go-webserver/routes"
	"github.com/jub3i/go-webserver/services/session"
	"log"
	"net/http"
	"os"
)

func main() {
	// config init
	c, err := config.Init([]string{
		"GOWS_PORT",
		"GOWS_SESSION_STORE_KEY",
		"GOWS_ENV",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("go-webserver mode: %s", c["GOWS_ENV"])

	// session init
	secure := false
	if os.Getenv("GOWS_ENV") == "prod" {
		secure = true
	}
	err = session.Init(
		"sid",
		os.Getenv("GOWS_SESSION_STORE_KEY"),
		secure,
	)
	if err != nil {
		log.Fatal(err)
	}

	// start the http server
	addr := "localhost:" + os.Getenv("GOWS_PORT")
	log.Printf("starting go-webserver on `%s`", addr)
	http.ListenAndServe(addr, routes.Mux)
}
