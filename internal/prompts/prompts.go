package prompts

import "github.com/charmbracelet/huh"

func AskQuestions() (*Config, error) {
	cfg := &Config{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project name").
				Value(&cfg.ProjectName),

			huh.NewSelect[string]().
				Title("Framework").
				Options(
					huh.NewOption("Gin", "gin"),
					huh.NewOption("Fiber", "fiber"),
				).
				Value(&cfg.Framework),

			huh.NewSelect[string]().
				Title("Database").
				Options(
					huh.NewOption("PostgreSQL", "postgres"),
					huh.NewOption("MySQL", "mysql"),
					huh.NewOption("None", "none"),
				).
				Value(&cfg.Database),

			huh.NewConfirm().
				Title("Enable JWT authentication?").
				Value(&cfg.JWT),

			huh.NewConfirm().
				Title("Add Docker support?").
				Value(&cfg.Docker),
		),
	)

	err := form.Run()

	return cfg, err
}
