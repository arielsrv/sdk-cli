package services

import (
	"os"

	"github.com/arielsrv/sdk-cli/pkg/model"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type GitService interface {
	Clone(template *model.Template) (*string, error)
}

type GitLabService struct {
	tempDirPattern string
}

func NewGitLabService() *GitLabService {
	return &GitLabService{
		tempDirPattern: "sdk-cli_",
	}
}

func (r GitLabService) Clone(template *model.Template) (*string, error) {
	tempDir, err := os.MkdirTemp("", r.tempDirPattern)
	if err != nil {
		return nil, err
	}

	plainClone, err := git.PlainClone(tempDir, false, &git.CloneOptions{
		URL:      template.RepositoryURL,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: "master_token",
			Password: os.Getenv("GITLAB_TOKEN"),
		},
	})

	if err != nil {
		return nil, err
	}

	tree, err := plainClone.Worktree()
	if err != nil {
		return nil, err
	}

	if err = tree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName("refs/tags/" + template.Tag),
	}); err != nil {
		return nil, err
	}

	return &tempDir, nil
}
