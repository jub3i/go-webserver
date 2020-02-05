package controllers

import (
	"fmt"
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

func Login() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "sid")
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
		}

		session.Values["authenticated"] = true
		session.Values["foo"] = "YEH"
		session.Values["bar"] = 123
		session.Save(r, w)
	}
	return http.HandlerFunc(fn)
}

func Logout() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "sid")
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
		}

		session.Values["authenticated"] = false
		session.Save(r, w)
	}
	return http.HandlerFunc(fn)
}

func Secret() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "sid")
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
		}

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		fmt.Println(session.Values)

		fmt.Fprintln(w, "My little Secret")
	}
	return http.HandlerFunc(fn)
}
