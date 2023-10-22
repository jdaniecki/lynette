//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Lint lynette source code
func Lint() error {
	return sh.Run("golangci-lint", "run")
}

// Build lynette binary
func Build() error {
	return sh.Run("go", "build", "-o", "./build/lynette", "cmd/lynette/lynette.go")
}

// Execute unit tests
func Test() error {
	mg.Deps(Build)
	env := map[string]string{
		"LYNETTE_BINARY_PATH": "../../build/lynette",
	}
	return sh.RunWith(env, "go", "test", "-v", "./...")
}

// Measure source code test coverage
func Coverage() error {
	env := map[string]string{
		"LYNETTE_BINARY_PATH": "../../build/lynette",
	}

	if err := sh.RunWith(env, "go", "test", "-v", "./...", "-cover", "-coverprofile=build/coverage.out", "-covermode=count"); err != nil {
		return err
	}
	if err := sh.Run("go", "tool", "cover", "-func=build/coverage.out"); err != nil {
		return err
	}
	if err := sh.Run("go", "tool", "cover", "-html=build/coverage.out", "-o=build/coverage.html"); err != nil {
		return err
	}
	return nil
}

// Delete build directory
func Clean() {
	os.RemoveAll("build")
}
