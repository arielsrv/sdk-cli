package model

type Template struct {
	Name          string `json:"name"`
	Language      string `json:"language"`
	RepositoryURL string `json:"repository_url"`
	Pattern       string `json:"pattern"`
}
