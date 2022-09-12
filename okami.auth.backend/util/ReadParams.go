package util

import (
	"net/http"
	"net/url"
	"strings"
)

// created at 06-20-2022
func GenerateQueryParam(request *http.Request) map[string]string {
	result := make(map[string]string)
	defer func() {
		_ = recover()
	}()
	var errs error
	rawQuery := request.URL.RawQuery
	rawSplit := strings.Split(rawQuery, "&")
	for key := range rawSplit {
		splitEqual := strings.Split(rawSplit[key], "=")
		result[splitEqual[0]], errs = url.QueryUnescape(splitEqual[1])
		if errs != nil {
			result[splitEqual[0]] = splitEqual[1]
		}
	}
	return result
}
