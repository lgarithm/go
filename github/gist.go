package github

import "time"

type GistFile struct {
	Content string `json:"content"`
}

type Gist struct {
	ID          string              `json:"id"`
	Description string              `json:"description"`
	Files       map[string]GistFile `json:"files"`
	UpdatedAt   time.Time           `json:"updated_at"`
}
