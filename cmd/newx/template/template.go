package template

import (
	"fmt"
	"github.com/fatih/color"
	"os"

	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/spf13/cobra"
)

type Command struct {
	Commands []*cobra.Command
}

func NewTemplateCommand(templateService services.TemplateService) *Command {
	templates := templateService.GetTemplates()
	commands := make([]*cobra.Command, 0, len(templates))
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
				fmt.Println()
				color.HiGreenString("Template %s created\n", appName)
				fmt.Println()
				fmt.Printf("  $ cd %s\n", appName)
				fmt.Println()
			},
		}
		templateCmd.PersistentFlags().StringVarP(&appName, "app-name", "", "", "an application name")
		err := templateCmd.MarkPersistentFlagRequired("app-name")
		if err != nil {
			templateCmd.PrintErr(err)
			os.Exit(1)
		}
		commands = append(commands, templateCmd)
	}

	return &Command{
		Commands: commands,
	}
}
