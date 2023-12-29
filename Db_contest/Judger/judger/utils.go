package judger

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemoveFilesWithSuffix(dir, suffix string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*."+suffix))
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			return err
		}
		fmt.Printf("Removed: %s\n", file)
	}

	return nil
}
