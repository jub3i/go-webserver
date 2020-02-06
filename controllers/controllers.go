package controllers

import (
	"net/http"
)

type Handler struct {
	Msg string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The Message Is: " + h.Msg))
}
