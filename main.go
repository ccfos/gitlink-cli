package main

import (
	"os"

	"github.com/gitlink-org/gitlink-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
