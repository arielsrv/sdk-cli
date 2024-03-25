package services

import (
	"cmp"
	_ "embed"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/arielsrv/sdk-cli/pkg/model"
	"github.com/samber/lo"
)

//go:embed template_service.json
var embeddedBytes []byte

var (
	templates []model.Template
)

func init() {
	err := json.Unmarshal(embeddedBytes, &templates)
	if err != nil {
		panic(err)
	}

	slices.SortFunc(templates, func(a, b model.Template) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
}

type TemplateService interface {
	GetTemplates() []model.Template
	GetAvailableLanguages() []model.Language
	GetTemplate(name string) (*model.Template, error)
	CreateTemplate(templateName string, appName string) error
}

type JSONTemplateService struct {
	gitlabService     GitService
	fileSystemService TreeService
}

func NewJSONTemplateService(
	gitlabService GitService,
	fileSystemService TreeService,
) *JSONTemplateService {
	return &JSONTemplateService{
		gitlabService:     gitlabService,
		fileSystemService: fileSystemService,
	}
}

func (r JSONTemplateService) GetTemplates() []model.Template {
	return templates
}

func (r JSONTemplateService) GetTemplate(name string) (*model.Template, error) {
	template, found := lo.Find(templates, func(item model.Template) bool {
		return item.ShortName == name
	})

	if !found {
		return nil, fmt.Errorf("template with key %s not found", name)
	}

	return &template, nil
}

func (r JSONTemplateService) GetAvailableLanguages() []model.Language {
	result := lo.Map(
		lo.Keys(
			lo.GroupBy(templates, func(item model.Template) string {
				return item.Language
			})), func(item string, index int) model.Language {
			return model.Language{
				Name: item,
			}
		})

	slices.SortFunc(result, func(a, b model.Language) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})

	return result
}

func (r JSONTemplateService) CreateTemplate(templateName string, appName string) error {
	template, err := r.GetTemplate(templateName)
	if err != nil {
		return err
	}

	path, err := r.gitlabService.Clone(template)
	if err != nil {
		return err
	}

	err = r.fileSystemService.WalkDir(*path, template.Pattern, appName)
	if err != nil {
		return err
	}

	return nil
}
