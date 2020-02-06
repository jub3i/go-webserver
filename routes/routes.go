package routes

import (
	c "github.com/jub3i/go-webserver/controllers"
	mw "github.com/jub3i/go-webserver/middleware"
	"github.com/justinas/alice"
	"net/http"
)

var Mux *http.ServeMux

func init() {
	Mux = http.NewServeMux()
	Mux.Handle("/api/auth/login", alice.New(mw.RequestLog).Then(c.Login()))
	Mux.Handle("/api/auth/logout", alice.New(mw.RequestLog).Then(c.Logout()))

	Mux.Handle("/secret", alice.New(mw.RequestLog, mw.IsAuthenticated).Then(c.Secret()))
}
