package services_test

import (
	"testing"

	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/stretchr/testify/assert"
)

func TestJSONServiceTemplate_GetTemplates(t *testing.T) {
	gitlabService := services.NewGitLabService()
	fileSystemService := services.NewFileSystemService()
	serviceTemplate := services.NewJSONTemplateService(gitlabService, fileSystemService)

	actual := serviceTemplate.GetTemplates()

	assert.GreaterOrEqual(t, len(actual), 1)
	assert.Equal(t, "Backend API SDK", actual[0].Name)
	assert.Equal(t, "backend-api-sdk", actual[0].ShortName)
	assert.Equal(t, "go", actual[0].Language)
	assert.Equal(t, "go-api", actual[0].Pattern)
	assert.Equal(t, "git@gitlab.com:iskaypetcom/digital/sre/tools/dev/backend-api-sdk.git", actual[0].RepositoryURL)
	assert.Equal(t, "v2.3.14", actual[0].Tag)
}

func TestJSONServiceTemplate_GetAvailableLanguages(t *testing.T) {
	gitlabService := services.NewGitLabService()
	fileSystemService := services.NewFileSystemService()
	serviceTemplate := services.NewJSONTemplateService(gitlabService, fileSystemService)
	actual := serviceTemplate.GetAvailableLanguages()

	assert.GreaterOrEqual(t, len(actual), 1)
	assert.Equal(t, "go", actual[0].Name)
}
