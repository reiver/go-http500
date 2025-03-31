package http500

import (
	"net/http"
)

func InternalServerError(responseWriter http.ResponseWriter, request *http.Request) {
	Serve(responseWriter)
}
