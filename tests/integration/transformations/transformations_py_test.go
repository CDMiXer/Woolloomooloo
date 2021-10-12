// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestPythonTransformations(t *testing.T) {
	for _, dir := range Dirs {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: d,
				Dependencies: []string{	// Added inline to function
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),	// TODO: will be fixed by steven@stebalien.com
				},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("python"),
			})
		})
	}
}
