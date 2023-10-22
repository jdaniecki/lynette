//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Lint lynette source code
func Lint() error {
	return sh.Run("golangci-lint", "run")
}

func buildGeneric() error {
	return sh.Run("go", "build", "-o", "./build/lynette", "cmd/lynette/lynette.go")
}

func buildCoverage() error {
	return sh.Run("go", "build", "-cover", "-o", "./build/lynette_coverage", "cmd/lynette/lynette.go")
}

// Build lynette binary
func Build() error {
	mg.Deps(buildGeneric, buildCoverage, downloadRootfs)
	return nil
}

// Downloads Ubuntu 22.04 base
func downloadRootfs() error {
	fsDir := filepath.Join("build", "rootfs")

	if _, exists := os.Stat(fsDir); exists == nil {
		fmt.Println("Skiping download as rootfs dir exists.")
		return nil
	}

	err := sh.Run("mkdir", "-p", fsDir)
	if err != nil {
		return err
	}

	fsFile := filepath.Join("build", "ubuntu.tar.gz")
	sh.Run("wget", "https://cdimage.ubuntu.com/ubuntu-base/releases/22.04/release/ubuntu-base-22.04-base-amd64.tar.gz", "-O", fsFile)
	if err != nil {
		return err
	}

	err = sh.Run("tar", "xvf", fsFile, "-C", fsDir)
	if err != nil {
		return err
	}
	return nil
}

// Execute unit tests
func Test() error {
	mg.SerialDeps(Build)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	buildDir := path.Join(wd, "build")
	env := map[string]string{
		"LYNETTE_BINARY_PATH": filepath.Join(buildDir, "lynette"),
		"ROOTFS":              filepath.Join(buildDir, "rootfs"),
	}
	return sh.RunWith(env, "go", "test", "-v", "./...")
}

// Measure source code test coverage
func Coverage() error {
	mg.SerialDeps(Clean, Build)

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	buildDir := path.Join(wd, "build")
	coverageDir := path.Join(buildDir)
	env := map[string]string{
		"LYNETTE_BINARY_PATH": path.Join(buildDir, "lynette_coverage"),
		"ROOTFS":              filepath.Join(buildDir, "rootfs"),
		"GOCOVERDIR":          coverageDir,
	}

	if err := sh.RunWith(env, "go", "test", "-v", "./..."); err != nil {
		return err
	}

	if err := sh.RunWith(env, "go", "tool", "covdata", "percent", "-i", coverageDir); err != nil {
		return err
	}

	coveragePath := filepath.Join(coverageDir, "coverage.out")
	if err := sh.RunWith(env, "go", "tool", "covdata", "textfmt", "-i", coverageDir, "-o", coveragePath); err != nil {
		return err
	}

	reportPath := filepath.Join(coverageDir, "coverage.html")
	if err := sh.Run("go", "tool", "cover", "-html", coveragePath, "-o", reportPath); err != nil {
		return err
	}
	return nil
}

// Delete build directory
func Clean() {
	os.RemoveAll("build")
}
