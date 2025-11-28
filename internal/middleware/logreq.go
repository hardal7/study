package middleware

import (
	logger "chat/internal/util"
	"net/http"
	"time"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Debug(r.Method + " " + r.URL.Path + " " + time.Since(start).String())
	}
}
