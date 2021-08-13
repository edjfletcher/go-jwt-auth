package go_jwt_auth

import (
	"firebase.google.com/go/v4/auth"
	"net/http"
)

type AuthorisationMiddlewareRoleCheck func(role string) bool

type AuthorisationMiddlewareOptions struct {
	RoleCheck AuthorisationMiddlewareRoleCheck
}

func AuthorisationMiddlewareFactory(options AuthorisationMiddlewareOptions) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, ok := r.Context().Value("token").(*auth.Token)
			if !ok {
				writeErrorMessage(w, 401, "authorisation middleware requires a token on context")
				return
			}

			role, ok := token.Claims["role"].(string)
			if ok && role == "" {
				writeErrorMessage(w, 401, "could not find a role on token")
				return
			}

			if !options.RoleCheck(role) {
				writeErrorMessage(w, 401, "role does not pass role check")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
