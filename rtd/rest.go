package rtd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"

	"github.com/lgarithm/go/net/http/probe"
)

func (c *Client) getJSON(u string, i interface{}) error {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return readJSON(resp.Body, &i)
}

func (c *Client) getText(u string) (string, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}
	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// Do overrides http.Client.Do
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	log.Printf("%s %s", req.Method, req.URL.String())
	return c.client.Do(withDefaultTrace(req))
}

func withDefaultTrace(req *http.Request) *http.Request {
	var (
		dns  probe.Interval
		dial probe.Interval
	)
	// TODO: check event order
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			log.Printf("DNSStart: %v", info)
			dns.Begin = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			log.Printf("DNSDone: %v", info)
			dns.End = time.Now()
		},
		ConnectStart: func(network, addr string) {
			log.Printf("ConnectStart: %s, %s", network, addr)
			dial.Begin = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			log.Printf("ConnectDone: %s, %s, err: %v", network, addr, err)
			dial.End = time.Now()
		},
		GotFirstResponseByte: func() {
			log.Printf("GotFirstResponseByte")
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	return req
}

var readJSON = _readJSON

// var readJSON = _readJSONVerbose

func _readJSON(r io.Reader, i interface{}) error {
	return json.NewDecoder(r).Decode(&i)
}

func _readJSONVerbose(r io.Reader, i interface{}) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Fprint(os.Stderr, "readJSON:\n")
	fmt.Printf("%s\n", string(bs))
	return json.Unmarshal(bs, &i)
}
