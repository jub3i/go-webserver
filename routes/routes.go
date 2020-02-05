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
	Mux.Handle("/bye", alice.New(mw.ReqLog).Then(c.Handler{Msg: "BYE"}))
	Mux.Handle("/hi", alice.New(mw.ReqLog).Then(c.Handler{Msg: "HI"}))
	Mux.Handle("/deny", alice.New(mw.ReqLog, mw.IsAdmin).Then(c.Handler{Msg: "DENY"}))
}
