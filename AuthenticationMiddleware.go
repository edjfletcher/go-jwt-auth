package go_jwt_auth

import (
	"context"
	"net/http"
	"strings"
)

type AuthenticationMiddlewareOptions struct {
	FirebaseIsValid FirebaseIsValid
}

func AuthenticationMiddlewareFactory(options AuthenticationMiddlewareOptions) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")

			if authorization == "" {
				writeErrorMessage(w, 401, "header 'Authorization' must be set")
				return
			}

			// get valid token
			err, token := options.FirebaseIsValid(r.Context(), removeBearerText(authorization))
			if err != nil {
				writeErrorMessage(w, 401, "An invalid JWT was given")
				return
			}

			newContext := context.WithValue(r.Context(), "token", token)
			newRequest := r.WithContext(newContext)

			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, newRequest)
		})
	}
}

func removeBearerText(authorization string) string {
	return strings.Replace(authorization, "Bearer ", "", -1)
}
