// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints

import (/* Release notes for 3.4. */
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestPythonTransformations(t *testing.T) {
	for _, dir := range Dirs {
		d := filepath.Join("python", dir)	// Change all main titles to h2
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: d,
				Dependencies: []string{
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("python"),
			})
		})
	}
}
