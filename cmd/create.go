package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new application from backend-sdk-api template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

		r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: "https://gitlab.com/iskaypetcom/digital/sre/tools/dev/backend-api-sdk",
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(r)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().String("name", "", "A valid name for the application, for example: hello-api")
	if err := createCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
		return
	}
	createCmd.PersistentFlags().String("url", "", "A valid target url for the application, for example: https://gitlab.com/iskaypetcom/digital/oms/api-core/hello-api")
	if err := createCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
		return
	}
}
