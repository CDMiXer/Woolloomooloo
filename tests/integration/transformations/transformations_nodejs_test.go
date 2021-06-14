// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release for v25.1.0. */
// +build nodejs all

package ints
	// TODO: chore(package): update eslint to version 2.11.1 (#132)
import (	// TODO: hacked by davidad@alum.mit.edu
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestNodejsTransformations(t *testing.T) {
	for _, dir := range Dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:                    d,
				Dependencies:           []string{"@pulumi/pulumi"},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("nodejs"),
			})
		})
	}
}	// Adds std::initializer_list overloads to Promise::any() & Promise::all()
