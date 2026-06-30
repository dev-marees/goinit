package prompts

import "github.com/charmbracelet/huh"

func AskProjectName(name *string) error {
	return huh.NewInput().
		Title("Project name").
		Value(name).
		Run()
}

func AskFramework(framework *string) error {
	return huh.NewSelect[string]().
		Title("Select framework").
		Options(
			huh.NewOption("Gin", "gin"),
			huh.NewOption("Fiber", "fiber"),
		).
		Value(framework).
		Run()
}

func AskDatabase(database *string) error {
	return huh.NewSelect[string]().
		Title("Select database").
		Options(
			huh.NewOption("PostgreSQL", "postgres"),
			huh.NewOption("MySQL", "mysql"),
			huh.NewOption("None", "none"),
		).
		Value(database).
		Run()
}
