package github

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type token struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

func load() (*token, error) {
	bs, err := ioutil.ReadFile(tokenFile())
	if err != nil {
		return nil, err
	}
	var t token
	if err := json.Unmarshal(bs, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

// SaveToken saves token file to default location
func SaveToken(bs []byte) error {
	var t token
	if err := json.Unmarshal(bs, &t); err != nil {
		return err
	}
	filename := tokenFile()
	os.MkdirAll(path.Dir(filename), os.ModePerm)
	return ioutil.WriteFile(filename, bs, os.ModePerm)
}

func tokenFile() string {
	return path.Join(os.Getenv("HOME"), ".github", "token.json")
}
