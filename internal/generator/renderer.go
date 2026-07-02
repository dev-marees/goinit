package generator

import (
	"bytes"
	"text/template"

	templatefiles "github.com/dev-marees/goinit/internal/templates"
)

func Render(templateName string, data any) ([]byte, error) {
	tmpl, err := template.ParseFS(
		templatefiles.FS,
		templateName,
	)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
