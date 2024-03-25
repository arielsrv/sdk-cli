package cmd

import (
	"github.com/arielsrv/sdk-cli/pkg/container"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

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

func init() {
	listCmd.AddCommand(languagesCmd)
}
