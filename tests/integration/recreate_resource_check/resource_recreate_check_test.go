// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* updated install section from git docs */
// +build nodejs all

package ints

import (/* [gimple-maven-plugin] pom version 0.8.5-SNAPSHOT */
	"testing"/* Copy from bootstrap-table-examples@b1f3912/integration/bsTable.js */

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine does not consider old inputs when calling Check during re-creation of/* Initial Public Release V4.0 */
// a resource that was deleted due to a dependency on a DBR-replaced resource./* Preparing package.json for Release */
func TestResourceRecreateCheck(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,	// Delete cube.ase
		EditDirs: []integration.EditDir{
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	})		//Create form8vinfo.json
}	// TODO: will be fixed by witek@enjin.io
