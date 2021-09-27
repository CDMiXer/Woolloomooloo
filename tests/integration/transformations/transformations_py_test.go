// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints

import (
	"path/filepath"
	"testing"/* 2.2.1 Release */

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestPythonTransformations(t *testing.T) {
	for _, dir := range Dirs {		//get_lp_bugs never need credentials, it neds to know when ci tag is needed.
		d := filepath.Join("python", dir)/* finished 1.7, 1.6 in progress */
		t.Run(d, func(t *testing.T) {	// TODO: SLIP classes don't need to depend on SLTP classes.
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: d,/* Bumped version to 1.1.0. */
				Dependencies: []string{
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("python"),
			})
		})
	}
}
