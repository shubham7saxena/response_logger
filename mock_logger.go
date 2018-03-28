package middlewares

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockResponseWriter struct {
	mock.Mock
}

func (m MockResponseWriter) Header() http.Header {
	return nil
}

func (m MockResponseWriter) Write(b []byte) (int, error) {
	return 0, nil
}

func (m MockResponseWriter) WriteHeader(statusCode int) {
	m.Called(statusCode)
}
