package httputil

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// FormFile gets a named file from http.Request
func FormFile(req *http.Request, name string) ([]byte, error) {
	f, _, err := req.FormFile(name)
	if err != nil {
		return nil, fmt.Errorf("%s not exist in FormFile: %v", name, err)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s from FormFile: %v", name, err)
	}
	return bs, nil
}
