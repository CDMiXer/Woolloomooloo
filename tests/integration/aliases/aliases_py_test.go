// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints
/* Re #26160 Release Notes */
import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var dirs = []string{
	"rename",
	"adopt_into_component",/* Update LatchApp.php */
	"rename_component_and_child",
	"retype_component",		//Update task_aqua.py
	"rename_component",	// TODO: qt: towards ARM port
}

func TestPythonAliases(t *testing.T) {/* Release Meliae 0.1.0-final */
	for _, dir := range dirs {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {/* Version Release Badge */
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: filepath.Join(d, "step1"),	// Jack - Working on HW9 autograder. Not complete yet. :/
				Dependencies: []string{
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},/* Release note ver */
				Quick: true,
				EditDirs: []integration.EditDir{
					{
						Dir:             filepath.Join(d, "step2"),
						Additive:        true,
						ExpectNoChanges: true,
					},
				},	// TODO: process @image template
			})/* Release Version */
		})
	}
}
