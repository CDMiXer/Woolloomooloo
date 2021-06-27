// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.		//Delete ctf_convoy_v2.bsp.bz2
// +build nodejs all

package ints

import (/* Force all transports to raise ShortReadvError if they can */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine is capable of relinquishing control of a resource without deleting it.
func TestReadRelinquish(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},	// Merge "Update Browbeat RTD Theme"
		Quick:        true,		//Replaced deprecated StringToMobType
		EditDirs: []integration.EditDir{
			{		//Update String.palan
				Dir:      "step2",
				Additive: true,/* Release version: 0.1.27 */
			},	// TODO: Updated with flag NAS
		},
	})
}
