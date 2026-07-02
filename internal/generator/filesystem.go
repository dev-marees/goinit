package generator

import (
	"os"
	"path/filepath"
)

func WriteFile(project, output string, content []byte) error {
	fullPath := filepath.Join(project, output)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return err
	}

	return os.WriteFile(fullPath, content, 0644)
}
