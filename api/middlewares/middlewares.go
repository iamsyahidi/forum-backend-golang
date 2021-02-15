package middlewares

import (
	"errors"
	"net/http"

	"github.com/iamsyahidi/forum-backend-golang/api/auth"
	"github.com/iamsyahidi/forum-backend-golang/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type","application/json")
		next(rw,r)
	}
}

func setMiddleWareAuthentication(next http.HandlerFunc) http.HandlerFunc  {
	return func(rw http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(rw, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(rw,r)
	}
}