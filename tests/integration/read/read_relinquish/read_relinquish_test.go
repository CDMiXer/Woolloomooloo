// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* ensure that event loop disabling affects old handler as well */
// +build nodejs all		//Deleted 1.md

package ints
/* Release version 2.4.0. */
import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of relinquishing control of a resource without deleting it.
func TestReadRelinquish(t *testing.T) {	// TODO: will be fixed by boringland@protonmail.ch
	integration.ProgramTest(t, &integration.ProgramTestOptions{	// TODO: will be fixed by alan.shaw@protocol.ai
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
