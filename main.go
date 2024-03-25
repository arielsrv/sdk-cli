package main

import (
	"os"

	"github.com/arielsrv/sdk-cli/cmd"
	"github.com/arielsrv/sdk-cli/pkg/container"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"go.uber.org/dig"
)

func main() {
	// services
	container.Inject(services.NewGitLabService, dig.As(new(services.GitService)))
	container.Inject(services.NewFileSystemService, dig.As(new(services.TreeService)))
	container.Inject(services.NewJSONTemplateService, dig.As(new(services.TemplateService)))

	// commands
	container.Inject(cmd.NewListCommand)
	container.Inject(cmd.NewNewCommand)
	container.Inject(cmd.NewTemplateCommand)

	// commands
	container.Inject(func() *cmd.RootCommand {
		rootCmd := cmd.NewRootCommand()
		listCmd := container.Provide[*cmd.ListCommand]()
		newCmd := container.Provide[*cmd.NewCommand]()
		templateCmd := container.Provide[*cmd.TemplateCommand]()
		for i := range templateCmd.Commands {
			newCmd.AddCommand(templateCmd.Commands[i])
		}

		rootCmd.AddCommand(listCmd.Command)
		rootCmd.AddCommand(newCmd.Command)

		return rootCmd
	})

	// execute
	rootCmd := container.Provide[*cmd.RootCommand]()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
