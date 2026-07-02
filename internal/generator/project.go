package generator

import "github.com/dev-marees/goinit/internal/prompts"

type File struct {
	Output   string
	Template string
}

func createProject(cfg prompts.Config) error {

	files := []File{
		{
			Output:   "go.mod",
			Template: "common/go.mod.tmpl",
		},
		{
			Output:   ".env",
			Template: "common/env.tmpl",
		},
		{
			Output:   ".gitignore",
			Template: "common/gitignore.tmpl",
		},
		{
			Output:   "README.md",
			Template: "common/readme.md.tmpl",
		},
	}

	switch cfg.Framework {
	case "gin":
		files = append(files, File{
			Output:   "cmd/server/main.go",
			Template: "gin/main.go.tmpl",
		})

	case "fiber":
		files = append(files, File{
			Output:   "cmd/server/main.go",
			Template: "fiber/main.go.tmpl",
		})
	}

	for _, file := range files {

		content, err := Render(file.Template, cfg)
		if err != nil {
			return err
		}

		if err := WriteFile(cfg.ProjectName, file.Output, content); err != nil {
			return err
		}
	}

	return nil
}
