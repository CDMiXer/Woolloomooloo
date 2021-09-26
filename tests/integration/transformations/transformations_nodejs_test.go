// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Merge "Add Service Graph documentation" */
// +build nodejs all

package ints	// TODO: exceptions to case_prep based on prepositional object
		//changes in the description
import (
	"path/filepath"
	"testing"
		//gigpress table experiment
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
				ExtraRuntimeValidation: Validator("nodejs"),/* set rectangle for display */
			})
		})/* Create to.mk */
	}
}
