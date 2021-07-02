package check

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Format checks golang's code format
func Format() error {
	if output, err := sh.Output("gofmt", "-e", "-d", "./.."); err != nil {
		return err
	} else if len(output) > 0 {
		return mg.Fatal(2, output)
	}
	return nil
}
