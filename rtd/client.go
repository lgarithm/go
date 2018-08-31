package rtd

import (
	"fmt"
	"net/http"
	"net/url"
)

// Client provides services defined at https://docs.readthedocs.io/en/latest/api/v2.html
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
	// TODO: implement
	return nil
}

// ListProject implements https://docs.readthedocs.io/en/latest/api/v2.html#project-list
// GET /api/v2/project/
func (c *Client) ListProject(slug string) ([]Project, error) {
	q := &url.Values{}
	q.Set("slug", slug)
	u := withQuery(withPath(c.endpoint, `project`), q)
	u.Path += "/"
	var result ListProjectResult
	if err := c.getJSON(u.String(), &result); err != nil {
		return nil, err
	}
	return result.Results, nil
}

// GetProject implements https://docs.readthedocs.io/en/latest/api/v2.html#project-details
// GET /api/v2/project/(int: id)/
func (c *Client) GetProject(id int) (*Project, error) {
	u := withPath(c.endpoint, fmt.Sprintf(`project/%d`, id))
	u.Path += "/"
	var project Project
	if err := c.getJSON(u.String(), &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// ListVersion implements https://docs.readthedocs.io/en/latest/api/v2.html#version-list
// GET /api/v2/version/
func (c *Client) ListVersion(slug string) ([]Version, error) {
	q := &url.Values{}
	q.Set("project__slug", slug)
	q.Set("active", "true")
	u := withQuery(withPath(c.endpoint, `version`), q)
	u.Path += "/"
	var result ListVersionResult
	if err := c.getJSON(u.String(), &result); err != nil {
		return nil, err
	}
	return result.Results, nil
}

// GetVersion implements https://docs.readthedocs.io/en/latest/api/v2.html#version-detail
// GET /api/v2/version/(int: id)/
func (c *Client) GetVersion(id int) (*Version, error) {
	u := withPath(c.endpoint, fmt.Sprintf(`version/%d`, id))
	u.Path += "/"
	var version Version
	if err := c.getJSON(u.String(), &version); err != nil {
		return nil, err
	}
	return &version, nil
}

// ListBuild implements https://docs.readthedocs.io/en/latest/api/v2.html#build-list
// GET /api/v2/build/
func (c *Client) ListBuild(slug string) ([]Build, error) {
	q := &url.Values{}
	q.Set("project__slug", slug)
	u := withQuery(withPath(c.endpoint, `build`), q)
	u.Path += "/"
	var result ListBuildResult
	if err := c.getJSON(u.String(), &result); err != nil {
		return nil, err
	}
	return result.Results, nil
}

// GetBuild implements https://docs.readthedocs.io/en/latest/api/v2.html#build-detail
// GET /api/v2/build/(int: id)/
func (c *Client) GetBuild(id int) (*Build, error) {
	u := withPath(c.endpoint, fmt.Sprintf(`build/%d`, id))
	u.Path += "/"
	var build Build
	if err := c.getJSON(u.String(), &build); err != nil {
		return nil, err
	}
	return &build, nil
}

// GetBuildLog returns the build log as txt
func (c *Client) GetBuildLog(id int) (string, error) {
	u := withPath(c.endpoint, fmt.Sprintf(`build/%d`, id))
	u.Path += ".txt"
	return c.getText(u.String())
}
