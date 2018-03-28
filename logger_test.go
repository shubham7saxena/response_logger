package middlewares

import (
	"testing"
)

func TestWriteHeader(t *testing.T) {
	statusCode := 200
	mrw := new(MockResponseWriter)
	mrw.On("WriteHeader", statusCode).Return().Once()
	lrw := newLoggingResponseWriter(mrw)
	lrw.WriteHeader(statusCode)
	mrw.AssertExpectations(t)
}
