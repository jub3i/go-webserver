package controllers

import (
	"log"
	"net/http"
)

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
		session.Options.MaxAge = -1
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

		log.Printf("%v\n", session.Values)
		w.Write([]byte("My little Secret"))
	}
	return http.HandlerFunc(fn)
}
