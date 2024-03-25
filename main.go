package main

import (
	"fmt"
	"os"

	"github.com/arielsrv/sdk-cli/cmd"
	"github.com/arielsrv/sdk-cli/pkg/container"
)

func main() {
	if err := container.Registry.Invoke(func(rootCmd *cmd.RootCommand) {
		if err := rootCmd.Execute(); err != nil {
			os.Exit(1)
		}
	}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
