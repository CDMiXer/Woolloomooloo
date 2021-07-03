// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all/* ff75ed7c-2e42-11e5-9284-b827eb9e62be */

package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* Create PackageUtil.java */
)

// Test that the engine does not consider old inputs when calling Check during re-creation of
// a resource that was deleted due to a dependency on a DBR-replaced resource./* Configure autoReleaseAfterClose */
func TestResourceRecreateCheck(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},/* Create Kaltura_API_Authentication_and_Security.md */
		Quick:        true,
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}		//Update default chart version to v0.1.1
