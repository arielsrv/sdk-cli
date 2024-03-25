package model

type Template struct {
	Name          string `json:"name"`
	ShortName     string `json:"short_name"`
	Description   string `json:"description"`
	Language      string `json:"language"`
	RepositoryURL string `json:"repository_url"`
	Pattern       string `json:"pattern"`
	Tag           string `json:"tag"`
}
