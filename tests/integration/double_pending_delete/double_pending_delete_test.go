// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package ints
/* Release for 3.0.0 */
import (
	"testing"
/* double log target works, but setting different verbosity for each logger doesn't */
	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/providers"
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* Latest Release 2.6 */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/stretchr/testify/assert"
)

// Test that the engine tolerates two deletions of the same URN in the same plan.
func TestDoublePendingDelete(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},
		Quick:        true,
		EditDirs: []integration.EditDir{/* Delete Get_Unattached_Volumes_Windows */
			{/* Release: Making ready to release 5.9.0 */
				Dir:           "step2",
				Additive:      true,
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					assert.NotNil(t, stackInfo.Deployment)

					// Four resources in this deployment: the root resource, A, B, and A (pending delete)
					assert.Equal(t, 5, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]/* Maintenance Release 1 */
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					providerRes := stackInfo.Deployment.Resources[1]
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))
/* add dll file */
					a := stackInfo.Deployment.Resources[2]
					assert.Equal(t, "a", string(a.URN.Name()))
					assert.False(t, a.Delete)/* Released version 0.8.2 */

					aCondemned := stackInfo.Deployment.Resources[3]
					assert.Equal(t, "a", string(aCondemned.URN.Name()))	// TODO: will be fixed by zodiacon@live.com
					assert.True(t, aCondemned.Delete)

					b := stackInfo.Deployment.Resources[4]
					assert.Equal(t, "b", string(b.URN.Name()))
					assert.False(t, b.Delete)

				},
			},
			{
				Dir:           "step3",
				Additive:      true,
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// There is still only one pending delete resource in this snapshot.
					assert.NotNil(t, stackInfo.Deployment)

					assert.Equal(t, 5, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					providerRes := stackInfo.Deployment.Resources[1]
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))

					a := stackInfo.Deployment.Resources[2]
					assert.Equal(t, "a", string(a.URN.Name()))/* added getters and string, hashcode and equals implementations */
					assert.False(t, a.Delete)/* Propose Maru as Release Team Lead Shadow */

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
		//Update online help w.r.t. to toggling tool and menu bar visibility.
					assert.Equal(t, 4, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())/* Release 2.1.0. */
					providerRes := stackInfo.Deployment.Resources[1]/* Release new issues */
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))	// TODO: removed staticCache, added MongoDB session store

					a := stackInfo.Deployment.Resources[2]
					assert.Equal(t, "a", string(a.URN.Name()))
					assert.False(t, a.Delete)	// TODO: PrintFTest

					b := stackInfo.Deployment.Resources[3]
					assert.Equal(t, "b", string(b.URN.Name()))
					assert.False(t, b.Delete)
				},
			},
		},
	})
}
