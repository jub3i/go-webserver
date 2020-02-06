package middleware

import (
	"context"
	s "github.com/jub3i/go-webserver/services/session"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func IsAuthenticated(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, err := s.Get(r)
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		// TODO: split this into two different checks and return 401 or 500
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized)
			return
		}

		// pass session data downstream using the request's context
		ctx := context.WithValue(r.Context(), "userid", 1337)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func RequestLog(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}
