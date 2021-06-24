// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"path/filepath"
	"testing"	// TODO: hacked by igor@soramitsu.co.jp

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)
	// move docs to top level
func TestNodejsTransformations(t *testing.T) {
	for _, dir := range Dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:                    d,
				Dependencies:           []string{"@pulumi/pulumi"},
				Quick:                  true,
				ExtraRuntimeValidation: Validator("nodejs"),	// isInstanceLoaded is obsolete
			})
		})		//New science page with some new content
	}
}
