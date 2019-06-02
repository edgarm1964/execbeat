package main

import (
	"os"

	"github.com/edgarm1964/execbeat/cmd"

	_ "github.com/edgarm1964/execbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
