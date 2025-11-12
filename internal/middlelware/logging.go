package middleware

import (
	"net/http"
	"time"
)

func ZeroLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Info().Str("method", r.Method).Str("path", r.URL.Path).Dur("dur", time.Since(start)).Msg("request")
	})
}
