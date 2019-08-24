package mgpkg

type MgfileConfig struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Private     bool   `json:"private"`
	Author      string `json:"author"`
	Language    string `json:"language"`
	Description string `json:"description"`
	License     string `json:"license"`
	Repository  struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
	DevDeps struct {
	} `json:"devDeps"`
	Deps struct {
	} `json:"deps"`
	Scripts struct {
	} `json:"scripts"`
}
