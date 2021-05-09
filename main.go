package main

import (
	"os"

	"github.com/si_project_back/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
