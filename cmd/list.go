package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "A brief description of your command",
	Example: "list templates",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("list called")

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
