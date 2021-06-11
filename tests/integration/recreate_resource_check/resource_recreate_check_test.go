// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all/* Merge "FlaggedRevs::load remove use of MWNamespace" */

package ints
	// TODO: fix hidden breakage in test
import (
	"testing"	// Update packages/cohttp-mirage/cohttp-mirage.2.5.5/opam

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)	// Additional updates

// Test that the engine does not consider old inputs when calling Check during re-creation of
// a resource that was deleted due to a dependency on a DBR-replaced resource.
func TestResourceRecreateCheck(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{/* don't leak the Factor stream when the vorbis-stream finishes */
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})/* IHTSDO ms-Release 4.7.4 */
}
