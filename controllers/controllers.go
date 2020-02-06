package controllers

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var key []byte
var path string
var store *sessions.FilesystemStore

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	key = []byte(os.Getenv("GOWS_SESSION_STORE_KEY"))
	path = filepath.Join(filepath.Dir(ex), "sessions")
	store = sessions.NewFilesystemStore(path, key)
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 14,
		HttpOnly: true,
		Secure:   true,
	}
}

type Handler struct {
	Msg string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The Message Is: " + h.Msg))
}
