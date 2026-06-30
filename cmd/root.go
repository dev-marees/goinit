package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:     "goinit",
	Short:   "GoInit is a Vite-like project generator for Go",
	Long:    "GoInit helps developers create production-ready Go projects.",
	Version: Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
