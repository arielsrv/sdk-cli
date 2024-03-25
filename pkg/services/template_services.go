package services

import "github.com/arielsrv/sdk-cli/pkg/model"

type TemplateService interface {
	GetTemplates() []model.Template
	GetAvailableLanguages() []model.Language
	GetTemplate(name string) (*model.Template, error)
	CreateTemplate(templateName string, appName string) error
}
