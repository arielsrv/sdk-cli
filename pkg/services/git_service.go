package services

import "github.com/arielsrv/sdk-cli/pkg/model"

type GitService interface {
	Clone(template *model.Template) (*string, error)
}
