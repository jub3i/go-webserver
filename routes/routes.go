package routes

import (
	"net/http"
	c "webserver/controllers"
	mw "webserver/middleware"
)

var Mux *http.ServeMux

func init() {
	Mux = http.NewServeMux()
	Mux.Handle("/bye", mw.ReqLog(c.Handler{Msg: "BYE"}))
	Mux.Handle("/hi", mw.ReqLog(c.Handler{Msg: "HI"}))
	Mux.Handle("/deny", mw.ReqLog(mw.IsAdmin(c.Handler{Msg: "DENY"})))
}
