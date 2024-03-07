package middleware

import (
	"log/slog"
	"net/http"
)

func LogStatus(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w2 := &responseWriter{ResponseWriter: w}
		h.ServeHTTP(w2, r)
		slog.Info("request", "method", r.Method, "url", r.URL.Path, "status", translateStatus(w2.status))
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func translateStatus(code int) int {
	if code == 0 {
		return http.StatusOK
	}
	return code
}
