package generator

import (
	"os"

	"github.com/dev-marees/goinit/internal/prompts"
)

func createFolders(cfg prompts.Config) error {
	folders := []string{
		cfg.ProjectName,
		cfg.ProjectName + "/cmd/server",
		cfg.ProjectName + "/internal/handler",
		cfg.ProjectName + "/internal/service",
		cfg.ProjectName + "/internal/repository",
		cfg.ProjectName + "/configs",
	}

	if cfg.Database != "none" {
		folders = append(
			folders,
			cfg.ProjectName+"/internal/database",
		)
	}

	for _, folder := range folders {
		if err := os.MkdirAll(folder, 0755); err != nil {
			return err
		}
	}

	return nil
}
