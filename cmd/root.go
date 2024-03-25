// Package cmd
// Copyright © 2024 IskayPet developers
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/spf13/cobra"
)

type RootCommand struct {
	*cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCmd := &cobra.Command{
		Use:   "sdk-cli",
		Short: "A CLI for IskayPet Apps",
	}
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return &RootCommand{
		Command: rootCmd,
	}
}

type ListCommand struct {
	*cobra.Command
}

func NewListCommand(templateService services.TemplateService) *ListCommand {
	return &ListCommand{
		Command: &cobra.Command{
			Use:   "list",
			Short: "List of available templates",
			Run: func(cmd *cobra.Command, args []string) {
				templateTable := table.New("Template Name", "Short Name", "Language", "Description").
					WithHeaderFormatter(color.New(color.FgGreen).SprintfFunc())

				templates := templateService.GetTemplates()
				for i := range templates {
					template := templates[i]
					templateTable.AddRow(template.Name, template.ShortName, template.Language, template.Description)
				}

				templateTable.Print()
			},
		},
	}
}

type NewCommand struct {
	*cobra.Command
}

func NewNewCommand(templateService services.TemplateService) *NewCommand {
	return &NewCommand{Command: &cobra.Command{
		Use:   "new",
		Short: "Creates a new app from short template name",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide a short template name")
				return
			}
			template, err := templateService.GetTemplate(args[0])
			if err != nil {
				fmt.Printf("Template %s not found\n", args[0])
				fmt.Printf("Please execute sdk-cli list to see available templates\n")
				return
			}
			fmt.Printf("Template %s found, creating ...\n", template.Name)
		},
	}}
}

type TemplateCommand struct {
	Commands []*cobra.Command
}

func NewTemplateCommand(templateService services.TemplateService) *TemplateCommand {
	templates := templateService.GetTemplates()
	var commands []*cobra.Command
	for i := range templates {
		template := templates[i]
		var appName string
		templateCmd := &cobra.Command{
			Use: template.ShortName,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Printf("Template %s found, creating ...\n", template.ShortName)
				err := templateService.CreateTemplate(template.ShortName, appName)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Printf("Template %s created\n", appName)
				fmt.Printf("Please execute cd %s\n", appName)
				fmt.Println()
			},
		}
		templateCmd.PersistentFlags().StringVarP(&appName, "app-name", "", "", "an application name")
		err := templateCmd.MarkPersistentFlagRequired("app-name")
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, templateCmd)
	}

	return &TemplateCommand{
		Commands: commands,
	}
}
