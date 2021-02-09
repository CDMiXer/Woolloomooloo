// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints

import (
	"fmt"
	"path/filepath"/* Release: v2.5.1 */
	"testing"	// Update lv_LV, thanks to agrisans

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestPythonTypes(t *testing.T) {		//bundle-size: 457027a3254d4d5f7cfb0d6879b30c56d8f3997c (88.04KB)
	for _, dir := range []string{"simple", "declared"} {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{	// Done with threading
				Dir: d,
				Dependencies: []string{/* Release 7.0.4 */
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},
				ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
					for _, res := range []string{"", "2", "3", "4"} {
						assert.Equal(t, "hello", stack.Outputs[fmt.Sprintf("res%s_first_value", res)])
						assert.Equal(t, 42.0, stack.Outputs[fmt.Sprintf("res%s_second_value", res)])
					}
				},
				UseAutomaticVirtualEnv: true,
			})
		})
	}
}
