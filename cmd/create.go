package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Welcome to GoInit")
	},
	Version: version,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
