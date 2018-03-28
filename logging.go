package middlewares

import (
	"net/http"

	"github.com/urfave/negroni"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) loggingResponseWriter {
	return loggingResponseWriter{w, http.StatusOK}
}

func (lrw loggingResponseWriter) WriteHeader(statusCode int) {
	lrw.statusCode = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (lrw loggingResponseWriter) StatusCode() int64 {
	return int64(lrw.statusCode)
}

func ResponseLogger() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		responseWriter := newLoggingResponseWriter(rw)
		next(responseWriter, r)
		// use response writer here
	})
}
