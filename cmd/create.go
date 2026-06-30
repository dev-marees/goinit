package cmd

import (
	"fmt"

	"github.com/dev-marees/goinit/internal/prompts"
	"github.com/spf13/cobra"
)

var (
	projectName string
	framework   string
	database    string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Go project",
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg := prompts.Config{
			ProjectName: projectName,
			Framework:   framework,
			Database:    database,
		}

		if cfg.ProjectName == "" {
			if err := prompts.AskProjectName(&cfg.ProjectName); err != nil {
				return err
			}
		}

		if cfg.Framework == "" {
			if err := prompts.AskFramework(&cfg.Framework); err != nil {
				return err
			}
		}

		if cfg.Database == "" {
			if err := prompts.AskDatabase(&cfg.Database); err != nil {
				return err
			}
		}

		fmt.Println()
		fmt.Println("Configuration")
		fmt.Println("-------------")
		fmt.Println("Project   :", cfg.ProjectName)
		fmt.Println("Framework :", cfg.Framework)
		fmt.Println("Database  :", cfg.Database)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(
		&projectName,
		"name",
		"n",
		"",
		"Project name",
	)

	createCmd.Flags().StringVar(
		&framework,
		"framework",
		"",
		"Framework (gin|fiber)",
	)

	createCmd.Flags().StringVar(
		&database,
		"database",
		"",
		"Database (postgres|mysql|none)",
	)
}
