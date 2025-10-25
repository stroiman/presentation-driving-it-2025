package main

import (
	"log/slog"
	"net/http"
)

func LogMiddleware(logger *slog.Logger, h http.Handler) http.Handler {
	if logger == nil {
		logger = slog.Default()
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("HTTP Request",
			slog.Group("request",
				"method", r.Method,
				"url", r.URL),
		)
		h.ServeHTTP(w, r)
	})
}
