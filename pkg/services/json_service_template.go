package services

import (
	"cmp"
	_ "embed"
	"encoding/json"
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

	slices.SortFunc(result, func(a, b model.Language) int {
		return cmp.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})

	return result
}
