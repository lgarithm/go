package travisci

import (
	"net/url"
	"path"
)

const endpoint = `https://api.travis-ci.com/`

var defaultEndpoint, _ = url.Parse(endpoint)

func withPath(base url.URL, p string) url.URL {
	u := base
	u.Path = path.Join(u.Path, p)
	return u
}

func withQuery(base url.URL, q *url.Values) url.URL {
	u := base
	u.RawQuery = q.Encode()
	return u
}
