package middleware

import (
	"net/http"
	"time"

	logger "github.com/hardal7/study/internal/util"
)

func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Debug(r.Method + " " + r.URL.Path + " " + time.Since(start).String())
	}
}
