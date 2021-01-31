package main

import (
	"net/http"
	"time"
)
import "github.com/sirupsen/logrus"

var logger *logrus.Entry

func init() {
	logger = logrus.WithFields(logrus.Fields{
		"component": "gowebsvc",
	})

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// Make a wrapper for ResonseWriter to be able to sava and later print the response code
type LoggingWriter struct {
	http.ResponseWriter // embedded writer type
	responseCode        int
}

func (lw *LoggingWriter) WriteHeader(h int) {
	lw.responseCode = h
	lw.ResponseWriter.WriteHeader(h)
}

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := &LoggingWriter{w, 0}
		inner.ServeHTTP(lw, r)

		logger.Debug(
			"%d\t%s\t%s\t%s\t%s",
			lw.responseCode,
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
