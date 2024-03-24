package services

import "github.com/spf13/myapp/pkg/model"

type GitService interface {
	Clone(template *model.Template) (*string, error)
}
