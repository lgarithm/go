package travisci

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

// Client provides services defined at https://docs.travis-ci.com/user/developer/#api-v3
type Client struct {
	client   *http.Client
	endpoint url.URL
}

// New creates new Client.
func New() *Client {
	return &Client{
		client:   &http.Client{},
		endpoint: *defaultEndpoint,
	}
}

// Ping checks the connectivity of the endpoint.
func (c *Client) Ping() error {
	u := c.endpoint
	u.Path = `/`
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	return nil
}
