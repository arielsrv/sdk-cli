package cmd

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/myapp/pkg/services"

	"github.com/spf13/cobra"
)

// languagesCmd represents the languages command
var languagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		table := table.New("Language").
			WithHeaderFormatter(color.New(color.FgGreen).SprintfFunc())

		languages := services.NewJSONServiceTemplate().GetAvailableLanguages()
		for _, template := range languages {
			table.AddRow(template.Name)
		}

		table.Print()
	},
}

func init() {
	listCmd.AddCommand(languagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// languagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// languagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
