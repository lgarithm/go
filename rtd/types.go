// Package rtd implements readthedocs API.
// https://docs.readthedocs.io/en/latest/api/v2.html
package rtd

// Project is https://docs.readthedocs.io/en/latest/api/v2.html#project-details
type Project struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Slug                string `json:"slug"`
	Description         string `json:"description"`
	Language            string `json:"language"`
	ProgrammingLanguage string `json:"programming_language"`
	Repo                string `json:"repo"`
	RepoType            string `json:"repo_type"`
	DefaultVersion      string `json:"default_version"`
	DefaultBranch       string `json:"default_branch"`
	DocumentationType   string `json:"documentation_type"`
	Users               []int  `json:"users"`
	CanonicalURL        string `json:"canonical_url"`
}

// Time is time in the format 2018-06-19T20:16:00.951959
type Time string

// BuildCommand contains the detail of a build step.
type BuildCommand struct {
	ID          int    `json:"id"`
	RunTime     int    `json:"run_time"`
	Command     string `json:"command"`
	Description string `json:"description"`
	Output      string `json:"output"`
	ExitCode    int    `json:"exit_code"`
	StartTime   Time   `json:"start_time"`
	EndTime     Time   `json:"end_time"`
	Build       int    `json:"build"`
}

// Build is https://docs.readthedocs.io/en/latest/api/v2.html#build-detail
type Build struct {
	ID           int            `json:"id"`
	Commands     []BuildCommand `json:"commands"`
	ProjectSlug  string         `json:"project_slug"`
	VersionSlug  string         `json:"version_slug"`
	DocsURL      string         `json:"docs_url"`
	StateDisplay string         `json:"state_display"`
	Type         string         `json:"type"`
	State        string         `json:"state"`
	Date         Time           `json:"date"`
	Success      bool           `json:"success"`
	Setup        string         `json:"setup"`
	SetupError   string         `json:"setup_error"`
	Output       string         `json:"output"`
	Error        string         `json:"error"`
	ExitCode     int            `json:"exit_code"`
	Commit       string         `json:"commit"`
	Length       int            `json:"length"`
	// cold_storage interface{} `json:"cold_storage"`
	Project int `json:"project"`
	Version int `json:"version"`
}

// Downloads is the set of all download links
type Downloads struct {
	PDF     string `json:"pdf"`
	HTMLzip string `json:"htmlzip"`
	Epub    string `json:"epub"`
}

// Version is https://docs.readthedocs.io/en/latest/api/v2.html#api-version-detail
type Version struct {
	ID          int       `json:"id"`
	Slug        string    `json:"slug"`
	VerboseName string    `json:"verbose_name"`
	Built       bool      `json:"built"`
	Active      bool      `json:"active"`
	Type        string    `json:"type"`
	Identifier  string    `json:"identifier"`
	Downloads   Downloads `json:"downloads"`
	Project     Project   `json:"project"`
}

// ListProjectResult is the structure returned by https://docs.readthedocs.io/en/latest/api/v2.html#project-list
type ListProjectResult struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Project `json:"results"`
}

// ListBuildResult is the structure returned by https://docs.readthedocs.io/en/latest/api/v2.html#build-list
type ListBuildResult struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Results  []Build `json:"results"`
}

// ListVersionResult is the structure returned by https://docs.readthedocs.io/en/latest/api/v2.html#version-list
type ListVersionResult struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Version `json:"results"`
}
