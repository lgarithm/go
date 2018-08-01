package probe

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
	"time"
)

// Probe is an HTTP probe
type Probe struct {
	url    string
	req    *http.Request
	client *http.Client
	tr     *http.Transport
}

// New create a new Probe with default settings
func New(target string, timeout time.Duration) (*Probe, error) {
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}
	tr := &http.Transport{}
	client := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}
	return &Probe{
		url:    target,
		req:    req,
		client: client,
		tr:     tr,
	}, nil
}

// Ping checks the target URL
func (p *Probe) Ping() {
	result := p.ping()
	fmt.Printf("%s %s\n", result, p.url)
}

func (p *Probe) ping() Result {
	var (
		dns  Interval
		dial Interval
	)
	// TODO: check event order
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dns.Begin = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			dns.End = time.Now()
		},
		ConnectStart: func(network, addr string) {
			dial.Begin = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			dial.End = time.Now()
		},
	}
	req := p.req.WithContext(httptrace.WithClientTrace(p.req.Context(), trace))
	var (
		rw        Interval
		recvErr   error
		recvBytes int
	)
	resp, err := p.client.Do(req)
	if err == nil {
		rw.Begin = time.Now()
		func() {
			defer resp.Body.Close()
			var bs []byte
			bs, recvErr = ioutil.ReadAll(resp.Body)
			recvBytes = len(bs)
		}()
		rw.End = time.Now()
	}
	// FIXME: clean DNS cache
	p.tr.CloseIdleConnections()
	return Result{
		DNSInterval:  dns,
		DialInterval: dial,
		RWInterval:   rw,

		RecvErr:   recvErr,
		RecvBytes: recvBytes,
	}
}
