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
		RepositoryURL: "https://gitlab.com/iskaypetcom/digital/sre/tools/dev/go-sdk-config",
		Tag:           "v0.0.9",
	})

	require.NoError(t, err)
	require.NotNil(t, path)

	treeService := services.NewFileSystemService()

	t.Logf("path: %s", *path)
	err = treeService.WalkDir(*path, "jqlang", appName)
	require.NoError(t, err)

	err = os.RemoveAll(*path)
	require.NoError(t, err)

	err = os.RemoveAll(appName)
	require.NoError(t, err)
}
