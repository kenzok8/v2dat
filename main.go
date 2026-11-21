package main

import (
	"os"

	"github.com/kenzok8/v2dat/cmd"
	_ "github.com/kenzok8/v2dat/cmd/unpack"
)

func main() {
	// cobra already prints the error; exit non-zero so callers can detect failure.
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
