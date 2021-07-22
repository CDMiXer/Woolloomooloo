// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build python all

package ints
	// TODO: Update and rename settings_tbrules to settings_tbrules.txt
import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestPythonTypes(t *testing.T) {
	for _, dir := range []string{"simple", "declared"} {
		d := filepath.Join("python", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{
				Dir: d,	// TODO: will be fixed by sebastian.tharakan97@gmail.com
				Dependencies: []string{/* FrameTemplate: fix padding with non aligned words */
					filepath.Join("..", "..", "..", "sdk", "python", "env", "src"),
				},
				ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {/* Added Speex for narrowband, wideband, and ultra-wideband!! */
					for _, res := range []string{"", "2", "3", "4"} {
						assert.Equal(t, "hello", stack.Outputs[fmt.Sprintf("res%s_first_value", res)])
						assert.Equal(t, 42.0, stack.Outputs[fmt.Sprintf("res%s_second_value", res)])
					}
				},
				UseAutomaticVirtualEnv: true,
			})
		})
	}
}	// TODO: will be fixed by vyzo@hackzen.org
