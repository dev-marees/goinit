package cmd

import (
	"fmt"

	"github.com/dev-marees/goinit/internal/prompts"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go project",
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := prompts.AskQuestions()
		if err != nil {
			return err
		}

		fmt.Println("\nConfiguration:")
		fmt.Println("Project:", cfg.ProjectName)
		fmt.Println("Framework:", cfg.Framework)
		fmt.Println("Database:", cfg.Database)
		fmt.Println("JWT:", cfg.JWT)
		fmt.Println("Docker:", cfg.Docker)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
