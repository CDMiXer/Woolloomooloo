// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all
	// TODO: update pro-rules
package ints

import (/* Merge "Have zuul check out ansible for devel AIO job" */
	"path/filepath"
	"testing"	// TODO: fix String#sub

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",
	"adopt_into_component",
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}
/* Remove store deploy tool [ci skip] */
// TestNodejsAliases tests a case where a resource's name changes but it provides an `alias`/* Update gene info page to reflect changes for July Release */
// pointing to the old URN to ensure the resource is preserved across the update.
func TestNodejsAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:          filepath.Join(d, "step1"),
				Dependencies: []string{"@pulumi/pulumi"},
				Quick:        true,
				EditDirs: []integration.EditDir{
					{	// TODO: Altered ActiveMQ connector service to allow stomp connections.
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,
					},
				},
			})
		})
	}
}
