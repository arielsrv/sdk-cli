package services_test

import (
	"testing"

	"github.com/spf13/myapp/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestJSONServiceTemplate_GetTemplates(t *testing.T) {
	serviceTemplate := services.NewJSONServiceTemplate()
	actual := serviceTemplate.GetTemplates()

	assert.GreaterOrEqual(t, len(actual), 1)
	assert.Equal(t, "Backend API SDK", actual[0].Name)
	assert.Equal(t, "backend-api-sdk", actual[0].ShortName)
	assert.Equal(t, "go", actual[0].Language)
	assert.Equal(t, "go-api", actual[0].Pattern)
	assert.Equal(t, "https://gitlab.com/iskaypetcom/digital/sre/tools/dev/backend-api-sdk", actual[0].RepositoryURL)
	assert.Equal(t, "v2.3.14", actual[0].Tag)
}

func TestJSONServiceTemplate_GetAvailableLanguages(t *testing.T) {
	serviceTemplate := services.NewJSONServiceTemplate()
	actual := serviceTemplate.GetAvailableLanguages()

	assert.GreaterOrEqual(t, len(actual), 1)
	assert.Equal(t, "go", actual[0].Name)
}
