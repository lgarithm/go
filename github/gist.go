package github

type GistFile struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string              `json:"description"`
	Files       map[string]GistFile `json:"files"`
}
