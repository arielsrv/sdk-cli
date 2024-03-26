package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/arielsrv/sdk-cli/cmd"
	"github.com/arielsrv/sdk-cli/pkg/container"
)

func main() {
	gitlabToken := os.Getenv("GITLAB_TOKEN")
	if gitlabToken == "" {
		fmt.Println(color.HiRedString("GITLAB_TOKEN is not set"))
		os.Exit(1)
	}

	if err := container.Registry.Invoke(func(rootCmd *cmd.RootCommand) {
		if err := rootCmd.Execute(); err != nil {
			os.Exit(1)
		}
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
