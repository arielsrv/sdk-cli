package newx

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/spf13/cobra"
)

type Command struct {
	*cobra.Command
}

func NewNewCommand(templateService services.TemplateService) *Command {
	return &Command{Command: &cobra.Command{
		Use:   "new",
		Short: "Creates a new app from short template name",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				color.HiRed("Please, provide a template name")
				fmt.Println()
				fmt.Println("  $ sdk-api list")
				fmt.Println()
				return
			}
			template, err := templateService.GetTemplate(args[0])
			if err != nil {
				color.HiRed(fmt.Sprintf("Template %s not found\n", args[0]))
				fmt.Println()
				fmt.Println("  $ sdk-api list")
				fmt.Println()
				return
			}
			fmt.Printf("Template for %s found, creating ...\n", template.Name)
		},
	}}
}
