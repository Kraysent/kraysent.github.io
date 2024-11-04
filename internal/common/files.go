package common

import (
	"os"
	"path/filepath"
)

func WriteFile(filePath, content string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), 0o770); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
