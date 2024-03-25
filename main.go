package main

import (
	"os"

	"github.com/arielsrv/sdk-cli/cmd"
	"github.com/arielsrv/sdk-cli/pkg/container"
)

func main() {
	rootCmd := container.Provide[*cmd.RootCommand]()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
