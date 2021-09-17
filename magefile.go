//go:build mage
// +build mage

// This file is part of template-terraform-module.
//
// template-terraform-module is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// template-terraform-module is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with template-terraform-module. If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	BuildDir      string = path.Join(".", "build")
	ExamplesDir   string = path.Join(".", "examples")
	LockTimeout   int    = 5
	ModuleDir     string = path.Join(".")
	PlanFilePath  string = path.Join(BuildDir, "module.plan")
	TestDir       string = path.Join(".", "test")
	InputVarsFile string = path.Join(ExamplesDir, "input.tfvars")
	VendorDir     string = path.Join(".", "vendor")
)

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS
// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// MAGE TARGETS
// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Build compiles the project for all the supported platforms
func Build() error {
	args := []string{
		"-chdir=" + ExamplesDir,
		"init",
		"-reconfigure",
	}

	if os.Getenv("CI") != "" {
		args = append(args, "-input=false", "-no-color")
	}

	err := sh.RunV("terraform", args...)

	if err != nil {
		return err
	}

	inputVarsFileAbsPath, err := filepath.Abs(InputVarsFile)

	if err != nil {
		return err
	}

	planFileAbsPath, err := filepath.Abs(PlanFilePath)

	if err != nil {
		return err
	}

	args = []string{
		"-chdir=" + ExamplesDir,
		"plan",
		"-var-file=" + inputVarsFileAbsPath,
		"-lock-timeout=" + strconv.Itoa(LockTimeout) + "s",
		"-out=" + planFileAbsPath,
	}

	if os.Getenv("CI") != "" {
		args = append(args, "-input=false", "-no-color")
	}

	err = sh.RunV("terraform", args...)

	if err != nil {
		return err
	}

	return nil
}

// Clean removes temporary and build files
func Clean() error {
	pathsToRemove := []string{
		path.Join(ExamplesDir, ".terraform"),
		path.Join(ExamplesDir, "terraform.tfstate"),
		path.Join(ExamplesDir, "terraform.tfstate.backup"),
		PlanFilePath,
	}

	for _, path := range pathsToRemove {
		if err := sh.Rm(path); err != nil {
			return err
		}
	}

	return nil
}

// Lint checks the project's code for style and syntax issues
func Lint() error {
	pathsToLint := []string{
		ExamplesDir,
		ModuleDir,
	}

	args := []string{
		"fmt",
		"-recursive",
	}

	if os.Getenv("CI") != "" {
		args = append(args, "-check", "-write=false", "-no-color")
	}

	for _, path := range pathsToLint {
		args = append(args, path)

		if err := sh.RunV("terraform", args...); err != nil {
			return err
		}

		args = args[:len(args)-1]
	}

	return nil
}

// Release generates a release tarball containing the built files for each supported platform
func Release() error {
	args := []string{
		"-chdir=" + ExamplesDir,
		"apply",
		"-lock-timeout=" + strconv.Itoa(LockTimeout) + "s",
	}

	if os.Getenv("CI") != "" {
		args = append(
			args,
			"-auto-approve",
			"-input=false",
			"-no-color",
		)
	}

	planFilePathAbs, err := filepath.Abs(PlanFilePath)

	if err != nil {
		return err
	}

	if _, err = os.Stat(planFilePathAbs); os.IsNotExist(err) {
		return errors.New("Plan file is not present, please run build before running release action!")
	}

	args = append(args, planFilePathAbs)

	err = sh.RunV("terraform", args...)

	if err != nil {
		return err
	}

	return nil
}

// Reset removes all files that Clean does plus the vendor directory
func Reset() error {
	mg.Deps(Clean)

	if err := sh.Rm(VendorDir); err != nil {
		return err
	}

	if err := os.Mkdir(VendorDir, 0755); err != nil {
		return err
	}

	return nil
}

// Scan runs a security check using Snyk to search for known vulnerabilities in project
func Scan() error {
	_, err := exec.LookPath("tfsec")

	if err != nil {
		return err
	}

	args := []string{
		ModuleDir,
		"--verbose",
		"--no-color",
	}

	return sh.RunV("tfsec", args...)
}

// Test runs the unit test for the project
func Test() error {
	args := []string{
		"test",
		"-v",
		"-count=1",
		"./" + TestDir + "/...",
	}

	return sh.RunV("go", args...)
}
