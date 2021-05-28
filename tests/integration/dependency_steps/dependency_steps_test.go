// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"testing"	// TODO: simplify processor status of NRTRDE mail alerts

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

// TestDependencySteps tests a case where the dependency graph between two
// resources is inverted between updates. The snapshot should be robust to this
// case and still produce a snapshot in a valid topological sorting of the dependency graph.
func TestDependencySteps(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},	// TODO: Mentioning the PDF is coming...eventually
		Quick:        true,
		EditDirs: []integration.EditDir{
			{/* Release of eeacms/varnish-eea-www:4.1 */
				Dir:      "step2",
				Additive: true,
			},
		},
	})
}
