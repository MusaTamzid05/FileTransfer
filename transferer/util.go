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

func GetSize(path string) (int64, error) {

	fPtr, err := os.Open(path)

	if err != nil {
		return -1, err
	}

	info, err := fPtr.Stat()

	if err != nil {
		return -1, err
	}

	return info.Size(), nil
}

func IsDir(path string) (bool, error) {

	file, err := os.Open(path)

	if err != nil {
		return false, err
	}

	defer file.Close()

	fi, err := file.Stat()

	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}
