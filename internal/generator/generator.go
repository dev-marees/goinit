package generator

import "github.com/dev-marees/goinit/internal/prompts"

func Generate(cfg prompts.Config) error {
	if err := createFolders(cfg); err != nil {
		return err
	}

	if err := createFiles(cfg); err != nil {
		return err
	}

	return nil
}
