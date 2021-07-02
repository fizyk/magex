package file

import (
	"errors"
	"os"
)

// Exists checks whether path exists within the system, or not
func Exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
