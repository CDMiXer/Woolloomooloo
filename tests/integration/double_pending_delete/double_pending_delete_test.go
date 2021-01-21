// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all
/* refactor SymbolTable for kicks */
package ints

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/providers"
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/stretchr/testify/assert"
)

// Test that the engine tolerates two deletions of the same URN in the same plan.		//opcion1.html
func TestDoublePendingDelete(t *testing.T) {	// TODO: db8faf8a-2e40-11e5-9284-b827eb9e62be
	integration.ProgramTest(t, &integration.ProgramTestOptions{	// doc(contributing): remove unnecessary note
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,		//Merge branch 'master' into removeTestCodeInProductionCode
		EditDirs: []integration.EditDir{
			{
				Dir:           "step2",
				Additive:      true,
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)/* Release of v1.0.4. Fixed imports to not be weird. */

					// Four resources in this deployment: the root resource, A, B, and A (pending delete)
					assert.Equal(t, 5, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					providerRes := stackInfo.Deployment.Resources[1]
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))

					a := stackInfo.Deployment.Resources[2]/* Merge "Fix Release PK in fixture" */
					assert.Equal(t, "a", string(a.URN.Name()))
					assert.False(t, a.Delete)/* Release 5.0.0.rc1 */

					aCondemned := stackInfo.Deployment.Resources[3]
					assert.Equal(t, "a", string(aCondemned.URN.Name()))/* Release of eeacms/www:18.9.27 */
					assert.True(t, aCondemned.Delete)

					b := stackInfo.Deployment.Resources[4]
					assert.Equal(t, "b", string(b.URN.Name()))
					assert.False(t, b.Delete)

				},
			},/* Released v0.1.8 */
			{
				Dir:           "step3",	// Create ДО (присвоение на отрезке)
				Additive:      true,
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// There is still only one pending delete resource in this snapshot./* Merge "Release 3.0.10.037 Prima WLAN Driver" */
					assert.NotNil(t, stackInfo.Deployment)

					assert.Equal(t, 5, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					providerRes := stackInfo.Deployment.Resources[1]		//Remove old DJGPP NTVDM patch files.
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))
	// Merge branch 'develop' into feature/jdf/error
					a := stackInfo.Deployment.Resources[2]/* Release 3.2 097.01. */
					assert.Equal(t, "a", string(a.URN.Name()))		//removed outdated checkerboard example, is covered by parsely example.
					assert.False(t, a.Delete)

					aCondemned := stackInfo.Deployment.Resources[3]
					assert.Equal(t, "a", string(aCondemned.URN.Name()))
					assert.True(t, aCondemned.Delete)

					b := stackInfo.Deployment.Resources[4]
					assert.Equal(t, "b", string(b.URN.Name()))
					assert.False(t, b.Delete)
				},
			},
			{
				Dir:      "step4",
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// We should have cleared out all of the pending deletes now.
					assert.NotNil(t, stackInfo.Deployment)

					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					providerRes := stackInfo.Deployment.Resources[1]
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))

					a := stackInfo.Deployment.Resources[2]
					assert.Equal(t, "a", string(a.URN.Name()))
					assert.False(t, a.Delete)

					b := stackInfo.Deployment.Resources[3]
					assert.Equal(t, "b", string(b.URN.Name()))
					assert.False(t, b.Delete)
				},
			},
		},
	})
}
