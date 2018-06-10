package httputil

import (
	"net/http"

	"github.com/lgarithm/go/net/http/contentdisposition"
)

// ServeDownload replies to the request with the content of the named file.
func ServeDownload(w http.ResponseWriter, r *http.Request, name string, content []byte) {
	contentdisposition.SetForFile(w.Header(), name)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(content)
}
