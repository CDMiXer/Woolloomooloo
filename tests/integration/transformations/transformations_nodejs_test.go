// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (	// TODO: will be fixed by souzau@yandex.com
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestNodejsTransformations(t *testing.T) {
	for _, dir := range Dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{/* Release v0.1.7 */
				Dir:                    d,/* DOC add link to Corda CLI UX Guide */
				Dependencies:           []string{"@pulumi/pulumi"},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("nodejs"),
			})
		})
	}
}
