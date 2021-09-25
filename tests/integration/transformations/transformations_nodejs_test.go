// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestNodejsTransformations(t *testing.T) {
	for _, dir := range Dirs {
		d := filepath.Join("nodejs", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir:                    d,	// TODO: hacked by nicksavers@gmail.com
				Dependencies:           []string{"@pulumi/pulumi"},/* Create Web.Release.config */
				Quick:                  true,/* Merge "[Release] Webkit2-efl-123997_0.11.109" into tizen_2.2 */
				ExtraRuntimeValidation: Validator("nodejs"),
			})/* An external DSL for Kendrick with all green tests */
		})
	}
}
