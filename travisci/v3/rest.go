package travisci

import (
	"log"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/lgarithm/go/net/http/probe"
)

// Do overrides http.Client.Do
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	log.Printf("%s %s", req.Method, req.URL.String())
	req.Header.Set("Travis-API-Version", "3")
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
