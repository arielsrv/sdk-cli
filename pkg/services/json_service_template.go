package services

import (
	"cmp"
	_ "embed"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"
	"github.com/spf13/myapp/pkg/model"
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

type JSONServiceTemplate struct {
	gitlabService     GitService
	fileSystemService TreeService
}

func NewJSONServiceTemplate(
	gitlabService GitService,
	fileSystemService TreeService,
) *JSONServiceTemplate {
	return &JSONServiceTemplate{
		gitlabService:     gitlabService,
		fileSystemService: fileSystemService,
	}
}

func (r JSONServiceTemplate) GetTemplates() []model.Template {
	return templates
}

func (r JSONServiceTemplate) GetTemplate(name string) (*model.Template, error) {
	template, found := lo.Find(templates, func(item model.Template) bool {
		return item.ShortName == name
	})

	if !found {
		return nil, fmt.Errorf("template with key %s not found", name)
	}

	return &template, nil
}

func (r JSONServiceTemplate) GetAvailableLanguages() []model.Language {
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

func (r JSONServiceTemplate) CreateTemplate(templateName string, appName string) error {
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
