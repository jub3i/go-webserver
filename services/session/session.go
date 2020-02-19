package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"path/filepath"
)

var store *sessions.FilesystemStore
var sessionName string

func Init(sn string, key string, secure bool) error {
	sessionName = sn

	ex, err := os.Executable()
	if err != nil {
		return err
	}
	path := filepath.Join(filepath.Dir(ex), "sessions")

	byteKey := []byte(key)
	store = sessions.NewFilesystemStore(path, byteKey)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 14,
		HttpOnly: true,
		Secure:   secure,
	}

	return nil
}

func Get(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, sessionName)
}
