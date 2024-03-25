// Package cmd
// Copyright © 2024 IskayPet developers
package cmd

import (
	"github.com/arielsrv/sdk-cli/pkg/container"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "sdk-cli",
	Short: "A CLI for IskayPet Apps",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of available templates",
	Run: func(cmd *cobra.Command, args []string) {
		table := table.New("Template Name", "Short Name", "Language", "Description").
			WithHeaderFormatter(color.New(color.FgGreen).SprintfFunc())

		serviceTemplate := container.Provide[services.TemplateService]()

		templates := serviceTemplate.GetTemplates()
		for i := range templates {
			template := templates[i]
			table.AddRow(template.Name, template.ShortName, template.Language, template.Description)
		}

		table.Print()
	},
}

// newCmd represents the new command.
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new app from template",
	Run: func(cmd *cobra.Command, args []string) {
		templateName, err := cmd.Flags().GetString("template")
		if err != nil {
			log.Fatal(err)
		}

		appName, err := cmd.Flags().GetString("app-name")
		if err != nil {
			log.Fatal(err)
		}

		serviceTemplate := container.Provide[services.TemplateService]()

		err = serviceTemplate.CreateTemplate(templateName, appName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	serviceTemplate := container.Provide[services.TemplateService]()
	templates := serviceTemplate.GetTemplates()
	for i := range templates {
		template := templates[i]
		templateCmd := &cobra.Command{
			Use:   template.ShortName,
			Short: template.Description,
			Run: func(cmd *cobra.Command, args []string) {
				appName, err := cmd.Flags().GetString("app-name")
				if err != nil {
					log.Fatal(err)
				}

				err = serviceTemplate.CreateTemplate(template.ShortName, appName)
				if err != nil {
					log.Fatal(err)
				}
			},
		}
		newCmd.PersistentFlags().String("app-name", "", "an app name to create from template and put in target dir")
		newCmd.AddCommand(templateCmd)
	}

	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(listCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
