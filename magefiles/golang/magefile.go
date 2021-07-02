package golang

import (
	magexTime "github.com/fizyk/magex/time"
	"github.com/magefile/mage/sh"
	"time"
)

// Format formats go code
func Format() error {
	magexTime.MeasureTime(time.Now(), "Go Format")
	return sh.RunV("go", "fmt", "./...")
}

// Tidy tidies go.mod http
func Tidy() error {
	magexTime.MeasureTime(time.Now(), "Tidy go.mod")
	return sh.RunV("go", "mod", "tidy")
}

// Test run tests for go code
func Test() error {
	magexTime.MeasureTime(time.Now(), "Run tests")
	return sh.RunV("go", "test", "-race", "-coverpkg", "./...", "-coverprofile=coverage.txt", "-covermode=atomic", "./...")
}
