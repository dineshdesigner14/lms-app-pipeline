package fileutil

import (
	"os"
)

func IsFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}
