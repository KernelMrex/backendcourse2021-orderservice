package transport

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var setContentTypeMiddleware = func(contentType string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		next.ServeHTTP(w, r)
	})
}

var logMiddleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(w, r)
		executionTime := time.Now().Sub(timeStart)
		log.WithFields(log.Fields{
			"methods": r.Method,
			"url": r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent": r.UserAgent(),
			"executionTime": time.Time{}.Add(executionTime).Format("15:04:05.99999999"),
		}).Info("got a new request")
	})
}
