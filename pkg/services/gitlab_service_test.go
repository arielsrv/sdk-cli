package services_test

import (
	"testing"

	"github.com/arielsrv/sdk-cli/pkg/model"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/stretchr/testify/require"
)

func TestGitLabService_Clone(t *testing.T) {
	gitLabService := services.NewGitLabService()

	path, err := gitLabService.Clone(&model.Template{
		RepositoryURL: "https://gitlab.com/iskaypetcom/digital/sre/tools/dev/go-sdk-config",
		Tag:           "v0.0.9",
	})

	require.NoError(t, err)
	require.NotNil(t, path)
}

func TestGitLabService_Clone_ReferenceNotFound(t *testing.T) {
	gitLabService := services.NewGitLabService()

	path, err := gitLabService.Clone(&model.Template{
		RepositoryURL: "https://gitlab.com/iskaypetcom/digital/sre/tools/dev/go-sdk-config",
		Tag:           "0.0.9",
	})

	require.Error(t, err)
	require.EqualError(t, err, "reference not found")
	require.Nil(t, path)
}
