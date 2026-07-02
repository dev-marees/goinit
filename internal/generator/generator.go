package generator

import "github.com/dev-marees/goinit/internal/prompts"

func Generate(cfg prompts.Config) error {
	return createProject(cfg)
}
