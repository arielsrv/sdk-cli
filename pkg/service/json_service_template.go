package service

import (
	_ "embed"
	"encoding/json"

	"github.com/samber/lo"
	"github.com/spf13/myapp/pkg/model"
)

//go:embed service_template.json
var embeddedBytes []byte

var (
	templates []model.Template
)

func init() {
	err := json.Unmarshal(embeddedBytes, &templates)
	if err != nil {
		panic(err)
	}
}

type JSONServiceTemplate struct {
}

func NewJSONServiceTemplate() *JSONServiceTemplate {
	return &JSONServiceTemplate{}
}

func (r JSONServiceTemplate) GetTemplates() []model.Template {
	return templates
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

	return result
}
