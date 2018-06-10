package restutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// ReplyJSON sends HTTP response as application/json
func ReplyJSON(w http.ResponseWriter, res interface{}) error {
	const applicationJSON = `application/json`
	w.Header().Set("Content-Type", applicationJSON)
	return WriteJSON(w, res)
}

// WriteJSON writes an interface in JSON format
func WriteJSON(w io.Writer, i interface{}) error {
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")
	return e.Encode(i)
}

// ReadJSON reads an JSON object
func ReadJSON(r io.Reader, i interface{}) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, i)
}

// GetJSON gets an JSON object from URL
func GetJSON(url string, i interface{}) error {
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}
	return ReadJSON(res.Body, i)
}

// PostJSON posts an JSON object to URL
func PostJSON(url string, i interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(i); err != nil {
		return err
	}
	client := http.Client{
		Timeout: 15 * time.Second,
	}
	_, err := client.Post(url, "application/json", &buf)
	return err
}
