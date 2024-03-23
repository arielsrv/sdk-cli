package services

import "github.com/spf13/myapp/pkg/model"

type TemplateService interface {
	GetTemplates() []model.Template
	GetAvailableLanguages() []model.Language
}
