package fileutil

import (
	"os"
	"path/filepath"
)

func RemoveAllFilesRecursively(folderPath string) error {
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if path == folderPath {
			return nil
		}
		if info.IsDir() {
			files, err := os.ReadDir(path)
			if err != nil {
				return nil
			}
			for _, file := range files {
				filePath := filepath.Join(path, file.Name())
				if _, err := os.Stat(filePath); err == nil || !os.IsNotExist(err) {
					err := os.Remove(filePath)
					if err != nil {
					} else {
					}
				}
			}

			err = os.Remove(path)
			if err != nil {
			} else {
			}
		} else {
			err := os.Remove(path)
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
			} else {
			}
		}
		return nil
	})
	return err
}
