package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/saufiroja/go-graphql-boilerplate/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		token := strings.Split(authorization, "Bearer ")[1]

		res, err := utils.ValidateToken(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "email", res)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
