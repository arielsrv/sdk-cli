package main

import (
	"github.com/spf13/myapp/cmd"
	"github.com/spf13/myapp/pkg/container"
	"github.com/spf13/myapp/pkg/services"
	"go.uber.org/dig"
)

func main() {
	container.Inject(services.NewGitLabService, dig.As(new(services.GitService)))
	container.Inject(services.NewFileSystemService, dig.As(new(services.TreeService)))
	container.Inject(services.NewJSONTemplateService, dig.As(new(services.TemplateService)))

	cmd.Execute()
}
