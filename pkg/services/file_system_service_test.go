package services_test

import (
	"os"
	"testing"

	"github.com/arielsrv/sdk-cli/pkg/model"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/stretchr/testify/require"
)

func TestFileSystemService_Walk(t *testing.T) {
	appName := "myapp"
	defer require.NoError(t, os.RemoveAll(appName))

	gitLabService := services.NewGitLabService()

	path, err := gitLabService.Clone(&model.Template{
		RepositoryURL: "https://github.com/jqlang/jq",
		Tag:           "jq-1.7.1",
	})

	require.NoError(t, err)
	require.NotNil(t, path)

	treeService := services.NewFileSystemService()

	err = treeService.WalkDir(*path, "jqlang", appName)
	require.NoError(t, err)
}
