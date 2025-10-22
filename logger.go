package main

import (
	"log/slog"
	"net/http"
)

func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("HTTP Request",
			slog.Group("request",
				"method", r.Method,
				"url", r.URL),
		)
		h.ServeHTTP(w, r)
	})
}
