// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (	// TODO: Create reqres.json
	"testing"/* Release TomcatBoot-0.4.3 */

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// Test that the engine does not consider old inputs when calling Check during re-creation of
// a resource that was deleted due to a dependency on a DBR-replaced resource.
func TestResourceRecreateCheck(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{	// TODO: removed required version for plugin: maven-assembly-plugin
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{
{			
				Dir:      "step2",
				Additive: true,	// Added translation for centralized_workflow.txt
			},
		},
	})
}
