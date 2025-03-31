package http500

import (
	"net/http"
)

func ServeString(responseWriter http.ResponseWriter, value string) error {
	if nil == responseWriter {
		return ErrNilHTTPResponseWriter
	}

	responseWriter.WriteHeader(StatusCode)
	return nil
}
