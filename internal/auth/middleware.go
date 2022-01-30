package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dineshdb/authnz/internal/utils"
	"github.com/rs/zerolog/log"
)

type key int

const (
	Scope   key = 0
	Subject key = 1
)

func (j *JWTValidator) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			log.Error().Msg("Invalid number of parts in headers")
			utils.Unauthorized(w)
			return
		}

		claims, err := j.Verify(authHeader[1])
		if err != nil {
			log.Error().Msg(fmt.Sprintf("%v", err))
			utils.Unauthorized(w)
			return
		}

		userId, err := strconv.Atoi(claims.Subject)
		if err != nil {
			utils.Unauthorized(w)
			log.Error().Msg("Invalid user id")
			return
		}

		r = r.WithContext(r.Context())
		r = r.WithContext(context.WithValue(r.Context(), Scope, claims.Scope))
		r = r.WithContext(context.WithValue(r.Context(), Subject, userId))
		next.ServeHTTP(w, r)
	})
}
