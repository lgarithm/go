package httputil

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

// NewMultipartRequest creates a multipart POST request
func NewMultipartRequest(url string, files map[string][]byte, fields map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range files {
		part, err := writer.CreateFormFile(k, k)
		if err != nil {
			return nil, err
		}
		if _, err := part.Write(v); err != nil {
			return nil, err
		}
	}
	for k, v := range fields {
		if err := writer.WriteField(k, v); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}
