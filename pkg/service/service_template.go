package service

import "github.com/spf13/myapp/pkg/model"

type Service interface {
	GetTemplates() []model.Template
	GetAvailableLanguages() []model.Language
}
