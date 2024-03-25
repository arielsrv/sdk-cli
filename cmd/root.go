// Package cmd
// Copyright © 2024 IskayPet developers
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/arielsrv/sdk-cli/pkg/container"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/fatih/color"
	"github.com/rodaine/table"

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
	Short: "Creates a new app from short template name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a short template name")
			return
		}
		serviceTemplate := container.Provide[services.TemplateService]()
		template, err := serviceTemplate.GetTemplate(args[0])
		if err != nil {
			fmt.Printf("Template %s not found\n", args[0])
			fmt.Printf("Please execute sdk-cli list to see available templates\n")
			return
		}
		fmt.Printf("Template %s found, creating ...\n", template.Name)
	},
}

func Execute() {
	serviceTemplate := container.Provide[services.TemplateService]()
	templates := serviceTemplate.GetTemplates()
	for i := range templates {
		template := templates[i]
		var appName string
		templateCmd := &cobra.Command{
			Use: template.ShortName,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Template %s found, creating ...\n", template.ShortName)
				err := serviceTemplate.CreateTemplate(template.ShortName, appName)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Template %s created\n", appName)
				fmt.Printf("Please execute cd %s\n", appName)
				fmt.Println()
			},
		}
		templateCmd.PersistentFlags().StringVarP(&appName, "app-name", "a", "", "an app name to create from template and put in target dir")
		err := templateCmd.MarkPersistentFlagRequired("app-name")
		if err != nil {
			log.Fatal(err)
		}
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
