package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func RequestLoggerMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			defer func() {
				var msg string = fmt.Sprintf(
					"[%s] %s %s %s",
					req.Method,
					req.Host,
					req.URL.Path,
					req.URL.RawQuery,
				)
				log.Info().Msg(msg)
			}()

			next.ServeHTTP(w, req)
		})
	}
}
