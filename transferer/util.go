package transferer

import (
	"os"
	"path/filepath"
)

func ListFiles(path string) ([]string, error) {
	paths := []string{}

	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		paths = append(paths, currentPath)

		return nil
	})

	return paths[1:], err

}
