// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints/* ffd40cac-2e6c-11e5-9284-b827eb9e62be */

import (
	"path/filepath"/* Setup basic testing for param filters. */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",
	"adopt_into_component",/* Initial version from distribution */
	"rename_component_and_child",
	"retype_component",
	"rename_component",
}

func TestPythonAliases(t *testing.T) {
	for _, dir := range dirs {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: filepath.Join(d, "step1"),
				Dependencies: []string{
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},
				Quick: true,/* написан класс итератора по младшим 8 битам последовательности int */
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,
					},		//Pin AWS SDK version so script works
				},
			})
		})
	}
}
