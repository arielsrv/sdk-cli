package cmd

import (
	"github.com/arielsrv/sdk-cli/pkg/container"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

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

func init() {
	listCmd.AddCommand(templatesCmd)
}
