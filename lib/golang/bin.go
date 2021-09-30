package golang

import (
	"fmt"
	"go/build"
	"os"
)

// BinPath returns Go's bin path
func BinPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return fmt.Sprintf("%s/bin", gopath)
}
