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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templatesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templatesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
