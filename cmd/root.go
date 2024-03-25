// Package cmd
// Copyright © 2024 IskayPet developers
package cmd

import (
	"github.com/arielsrv/sdk-cli/cmd/list"
	"github.com/arielsrv/sdk-cli/cmd/newx"
	"github.com/arielsrv/sdk-cli/cmd/newx/template"
	"github.com/arielsrv/sdk-cli/pkg/container"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
)

var LatestVersion = "0.0.1"

type RootCommand struct {
	*cobra.Command
}

func NewRootCommand(
	listCmd *list.Command,
	newCmd *newx.Command,
	templateCmd *template.Command,
) *RootCommand {
	rootCmd := &cobra.Command{
		Use:     "sdk-cli",
		Short:   color.HiYellowString("A CLI for IskayPet Apps"),
		Version: LatestVersion,
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	for i := range templateCmd.Commands {
		newCmd.AddCommand(templateCmd.Commands[i])
	}

	rootCmd.AddCommand(listCmd.Command)
	rootCmd.AddCommand(newCmd.Command)

	return &RootCommand{
		Command: rootCmd,
	}
}

func init() {
	// services
	container.Inject(services.NewGitLabService, dig.As(new(services.GitService)))
	container.Inject(services.NewFileSystemService, dig.As(new(services.TreeService)))
	container.Inject(services.NewJSONTemplateService, dig.As(new(services.TemplateService)))

	// commands
	container.Inject(list.NewListCommand)
	container.Inject(newx.NewNewCommand)
	container.Inject(template.NewTemplateCommand)
	container.Inject(NewRootCommand)
}
