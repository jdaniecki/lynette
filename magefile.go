//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/sh"
)

// Build lynette binary
func Build() error {
	return sh.Run("go", "build", "-o", "build", "cmd/lynette/lynette.go")
}

// Execute unit tests
func Test() error {
	return sh.Run("go", "test", "-v", "./...")
}

// Delete build directory
func Clean() {
	os.RemoveAll("build")
}
