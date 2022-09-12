package util

import (
	"mime/multipart"
	"net/http"
)

func ReadMultipartFile(request *http.Request, key string) (file multipart.File, filename string, size int64, errs error) {
	file, handler, errs := request.FormFile(key)
	if errs != nil {
		return
	}

	defer func() {
		errs = file.Close()
		if errs != nil {
			return
		}
	}()

	return file, handler.Filename, handler.Size, nil
}
