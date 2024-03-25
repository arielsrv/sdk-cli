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
	Short: "A brief description of your application",
}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "A brief description of your command",
	Example: "list templates",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("list called")

	},
}

// templatesCmd represents the templates command.
var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		table := table.New("Template Name", "Short Name", "Language").
			WithHeaderFormatter(color.New(color.FgGreen).SprintfFunc())

		serviceTemplate := container.Provide[services.TemplateService]()

		templates := serviceTemplate.GetTemplates()
		for _, template := range templates {
			table.AddRow(template.Name, template.ShortName, template.Language)
		}

		table.Print()
	},
}

// languagesCmd represents the languages command.
var languagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		table := table.New("Language").
			WithHeaderFormatter(color.New(color.FgGreen).SprintfFunc())

		serviceTemplate := container.Provide[services.TemplateService]()

		languages := serviceTemplate.GetAvailableLanguages()
		for _, template := range languages {
			table.AddRow(template.Name)
		}

		table.Print()
	},
}

// newCmd represents the new command.
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
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
	newCmd.PersistentFlags().String("app-name", "", "an app name to create from template and put in target dir")
	newCmd.PersistentFlags().String("template", "", "execute sdk-cli list templates for available templates")

	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(templatesCmd)
	listCmd.AddCommand(languagesCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
