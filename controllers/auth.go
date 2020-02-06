package controllers

import (
	s "github.com/jub3i/go-webserver/services/session"
	"log"
	"net/http"
)

func Login() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, err := s.Get(r)
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
		}

		session.Values["authenticated"] = true
		session.Save(r, w)
	}
	return http.HandlerFunc(fn)
}

func Logout() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, err := s.Get(r)
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
		session, err := s.Get(r)
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
		userid := r.Context().Value("userid")
		log.Printf("%v", userid)

		w.Write([]byte("My little Secret"))
	}
	return http.HandlerFunc(fn)
}
