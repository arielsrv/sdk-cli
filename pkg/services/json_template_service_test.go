package services_test

import (
	"testing"

	mocks "github.com/arielsrv/sdk-cli/pkg/mocks/pkg/services"
	"github.com/arielsrv/sdk-cli/pkg/model"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSONServiceTemplate_GetTemplates(t *testing.T) {
	gitlabService := mocks.NewMockGitService(t)
	fileSystemService := mocks.NewMockTreeService(t)
	serviceTemplate := services.NewJSONTemplateService(gitlabService, fileSystemService)

	actual := serviceTemplate.GetTemplates()

	assert.Len(t, actual, 1)
	assert.Equal(t, "Golang API", actual[0].Name)
	assert.Equal(t, "go-api", actual[0].ShortName)
	assert.Equal(t, "go", actual[0].Language)
	assert.Equal(t, "go-api", actual[0].Pattern)
	assert.Equal(t, "git@gitlab.com:iskaypetcom/digital/sre/tools/dev/backend-api-sdk.git", actual[0].RepositoryURL)
	assert.Equal(t, "v2.3.14", actual[0].Tag)
}

func TestJSONServiceTemplate_GetAvailableLanguages(t *testing.T) {
	gitlabService := mocks.NewMockGitService(t)
	fileSystemService := mocks.NewMockTreeService(t)
	serviceTemplate := services.NewJSONTemplateService(gitlabService, fileSystemService)
	actual := serviceTemplate.GetAvailableLanguages()

	assert.GreaterOrEqual(t, len(actual), 1)
	assert.Equal(t, "go", actual[0].Name)
}

func TestJSONTemplateService_CreateTemplate(t *testing.T) {
	gitlabService := mocks.NewMockGitService(t)
	path := "/path/to/template"
	gitlabService.EXPECT().Clone(&model.Template{
		Name:          "Golang API",
		ShortName:     "go-api",
		Description:   "Golang API from Backend API SDK w/kubernetes, docker and prometheus monitoring",
		Language:      "go",
		RepositoryURL: "git@gitlab.com:iskaypetcom/digital/sre/tools/dev/backend-api-sdk.git",
		Pattern:       "go-api",
		Tag:           "v2.3.14",
	}).Return(&path, nil)

	fileSystemService := mocks.NewMockTreeService(t)
	fileSystemService.EXPECT().WalkDir(path, "go-api", "hello-api").Return(nil)

	serviceTemplate := services.NewJSONTemplateService(gitlabService, fileSystemService)

	err := serviceTemplate.CreateTemplate("go-api", "hello-api")
	require.NoError(t, err)
}
