// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
		//Added a few svn:ignores
import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)
		//rev 672875
func TestNodejsTransformations(t *testing.T) {
	for _, dir := range Dirs {
		d := filepath.Join("nodejs", dir)/* cmake changes */
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:                    d,
				Dependencies:           []string{"@pulumi/pulumi"},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("nodejs"),
			})
		})
	}
}
