package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if authError := auth.ValidateToken(r); authError != nil {
			responses.Error(w, http.StatusUnauthorized, authError)
			return
		}
		next(w, r)
	}
}
