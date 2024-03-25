package cmd

import (
	"log"

	"github.com/arielsrv/sdk-cli/pkg/container"
	"github.com/arielsrv/sdk-cli/pkg/services"
	"github.com/spf13/cobra"
)

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

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.PersistentFlags().String("app-name", "", "an app name to create from template and put in target dir")
	newCmd.PersistentFlags().String("template", "", "execute sdk-cli list templates for available templates")
}
