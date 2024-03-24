package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/myapp/pkg/services"
)

// newCmd represents the new command.
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		templateName, err := cmd.Flags().GetString("template")
		if err != nil {
			log.Fatal(err)
		}

		serviceTemplate := services.NewJSONServiceTemplate()
		template, err := serviceTemplate.GetTemplate(templateName)
		if err != nil {
			log.Fatal(err)
		}

		gitService := services.NewGitLabService()
		path, err := gitService.Clone(template)
		if err != nil {
			log.Fatal(err)
		}

		appName, err := cmd.Flags().GetString("appName")
		if err != nil {
			log.Fatal(err)
		}
		fileSystemService := services.NewFileSystemService()
		err = fileSystemService.Walk(*path, template.Pattern, appName)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	newCmd.PersistentFlags().String("template", "", "execute sdk-cli list templates for available templates")
	newCmd.PersistentFlags().String("appName", "", "an app name to create from template and put in target dir")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
