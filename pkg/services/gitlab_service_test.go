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
		RepositoryURL: "https://github.com/jqlang/jq",
		Tag:           "jq-1.7.1",
	})

	require.NoError(t, err)
	require.NotNil(t, path)
}

func TestGitLabService_Clone_ErrAuthRequired(t *testing.T) {
	gitLabService := services.NewGitLabService()

	path, err := gitLabService.Clone(&model.Template{
		RepositoryURL: "https://github.com/jqlang/_not_found_",
		Tag:           "v2.3.11",
	})

	require.Error(t, err)
	require.EqualError(t, err, "authentication required")
	require.Nil(t, path)
}

func TestGitLabService_Clone_ReferenceNotFound(t *testing.T) {
	gitLabService := services.NewGitLabService()

	path, err := gitLabService.Clone(&model.Template{
		RepositoryURL: "https://github.com/jqlang/jq",
		Tag:           "jq-222.333.444",
	})

	require.Error(t, err)
	require.EqualError(t, err, "reference not found")
	require.Nil(t, path)
}
