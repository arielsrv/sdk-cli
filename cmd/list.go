package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
