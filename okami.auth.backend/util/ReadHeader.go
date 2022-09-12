package util

import (
	"net/http"
	"strings"
)

//created at 08-18-2022
func ReadHeader(request *http.Request, headerName string) (result string) {
	result = request.Header.Get(headerName)
	if result == "" {
		result = request.Header.Get(strings.ToLower(headerName))
	}
	return result
}
