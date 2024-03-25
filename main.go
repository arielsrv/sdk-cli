package main

import (
	"github.com/arielsrv/sdk-cli/cmd"
	"github.com/arielsrv/sdk-cli/pkg/container"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"go.uber.org/dig"
)

func main() {
	// commands

	// services
	container.Inject(services.NewGitLabService, dig.As(new(services.GitService)))
	container.Inject(services.NewFileSystemService, dig.As(new(services.TreeService)))
	container.Inject(services.NewJSONTemplateService, dig.As(new(services.TemplateService)))

	cmd.Execute()
}
