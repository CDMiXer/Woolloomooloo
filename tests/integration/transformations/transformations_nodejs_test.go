// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
	// TODO: [objc-arc] Added an option to arc-annotations for turning off CheckForCFGHazard.
import (
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
				Dependencies:           []string{"@pulumi/pulumi"},/* Pin websocket-client to latest version 0.57.0 */
				Quick:                  true,
				ExtraRuntimeValidation: Validator("nodejs"),
			})
		})
	}	// 72dc000f-2eae-11e5-890b-7831c1d44c14
}
