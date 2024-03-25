package list

import (
	"fmt"

	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type Command struct {
	*cobra.Command
}

func NewListCommand(templateService services.TemplateService) *Command {
	return &Command{
		Command: &cobra.Command{
			Use:   "list",
			Short: "List of available templates",
			Run: func(cmd *cobra.Command, args []string) {
				templateTable := table.New("Template Name", "Short Name", "Language", "Version", "Description").
					WithHeaderFormatter(color.New(color.FgGreen).SprintfFunc())

				templates := templateService.GetTemplates()
				for i := range templates {
					template := templates[i]
					templateTable.AddRow(template.Name, template.ShortName, template.Language, template.Tag, template.Description)
				}

				templateTable.Print()

				if len(templates) > 0 {
					fmt.Println()
					fmt.Println("Example:")
					fmt.Println()
					fmt.Printf("  $ sdk-api new %s\n", templates[0].ShortName)
					fmt.Println()
				}
			},
		},
	}
}
