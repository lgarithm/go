// Package github provides a simple github client which is dependency free.
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var (
	endpoint, _ = url.Parse(`https://api.github.com`)
)

type SimpleClient struct {
	endpoint *url.URL
}

func NewSimpleClient() (*SimpleClient, error) {
	t, err := load()
	if err != nil {
		return nil, err
	}
	u := endpoint
	u.User = url.UserPassword(t.User, t.Token)
	return &SimpleClient{
		endpoint: u,
	}, nil
}

func (c SimpleClient) CommentIssue(owner, repo string, id int, msg string) error {
	u := c.endpoint
	u.Path = fmt.Sprintf("/repos/%s/%s/issues/%d/comments", owner, repo, id)
	bs, _ := json.Marshal(struct {
		Body string `json:"body"`
	}{Body: msg})
	body := &bytes.Buffer{}
	body.Write(bs)
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func (c SimpleClient) CreateGist(desc string, files map[string]string) error {
	u := c.endpoint
	u.Path = "/gists"
	gistFiles := map[string]GistFile{}
	for name, content := range files {
		gistFiles[name] = GistFile{Content: content}
	}
	gist := Gist{
		Description: desc,
		Files:       gistFiles,
	}
	bs, _ := json.Marshal(gist)
	body := &bytes.Buffer{}
	body.Write(bs)
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return err
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
