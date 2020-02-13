package session

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var store *sessions.FilesystemStore
var sessionName = "sid"

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(filepath.Dir(ex), "sessions")

	secure := false
	if os.Getenv("GOWS_ENV") == "prod" {
		secure = true
	}

	key := []byte(os.Getenv("GOWS_SESSION_STORE_KEY"))
	store = sessions.NewFilesystemStore(path, key)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 14,
		HttpOnly: true,
		Secure:   secure,
	}
}

func Get(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, sessionName)
}
