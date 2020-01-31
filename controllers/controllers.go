package controllers

import (
	"fmt"
	"net/http"
)

type Handler struct {
	Msg string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The Message Is: "+h.Msg)
}

// var Handlerr = handler{msg: "ROOT"}
