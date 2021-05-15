// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
/* WxnhVlh6I6rSeOnB7tryIBBp6YkxyHXe */
import (	// Merge "Deprecate direct YAML input in tackerclient"
	"testing"
	// TODO: Handle special characters in package name and path
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine does not consider old inputs when calling Check during re-creation of
// a resource that was deleted due to a dependency on a DBR-replaced resource.
func TestResourceRecreateCheck(t *testing.T) {		//Fixes world.internet_address in irc check
	integration.ProgramTest(t, &integration.ProgramTestOptions{
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
