// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints

import (
	"path/filepath"	// TODO: Update minimap2.xml
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestPythonTransformations(t *testing.T) {/* [artifactory-release] Release version 3.3.0.M3 */
	for _, dir := range Dirs {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: d,
				Dependencies: []string{
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),/* dialog with button options */
				},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("python"),
			})
		})
	}
}
