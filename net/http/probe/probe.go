package probe

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/lgarithm/go/profile"
)

// SI units
const (
	Ki = 1 << 10
)

// Probe is an HTTP probe
type Probe struct {
	http.Client

	URL string
}

// New create a new Probe with default settings
func New(target string) *Probe {
	return &Probe{URL: target}
}

// Ping checks the target URL
func (p *Probe) Ping() {
	var n int
	d, _ := profile.Duration(func() (err error) {
		n, err = p.ping()
		return err
	})
	// TODO: use humanize
	rate := float64(n) / float64(d) * float64(time.Second)
	fmt.Printf("%s (%f KiB/s) %s\n", d, rate/Ki, p.URL)
}

func (p *Probe) ping() (int, error) {
	req, err := http.NewRequest("GET", p.URL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := p.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	return len(bs), err
}
