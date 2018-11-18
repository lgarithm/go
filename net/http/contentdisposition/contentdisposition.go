package contentdisposition

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Filename parse the filename from Content-Disposition
// val is assumed to has the form
// attachment; filename="src.tar.bz2"
func Filename(val string) string {
	parts := strings.Split(val, ";")
	if len(parts) != 2 {
		return ""
	}
	if parts[0] != "attachment" {
		return ""
	}
	parts[1] = strings.Trim(parts[1], " \t")
	parts = strings.Split(parts[1], "=")
	if len(parts) != 2 {
		return ""
	}
	if parts[0] != "filename" {
		return ""
	}
	val, err := strconv.Unquote(parts[1])
	if err != nil {
		return ""
	}
	return val
}

// SetForFile sets the Content-Disposition header for given filename.
func SetForFile(header http.Header, filename string) {
	header.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
}
