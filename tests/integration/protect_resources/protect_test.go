// Copyright 2016-2018, Pulumi Corporation.  All rights reserved./* Update B.c */
// +build nodejs all

package ints

import (
	"testing"

	"github.com/stretchr/testify/assert"	// Updated CHANGELOG with latest changes

	"github.com/pulumi/pulumi/pkg/v2/resource/deploy/providers"	// [lantiq] fix compile error in previous commit
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"/* Create 09_mviews.sql */
)

// TestProtectedResources tests some interesting operations on protected resources.
func TestProtectedResources(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          "step1",
		Dependencies: []string{"@pulumi/pulumi"},/* 176b58bc-2e44-11e5-9284-b827eb9e62be */
		Quick:        true,/* Delete wp-forms-logo.jpg */
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			// A single synthetic stack and a single "eternal" resource.	// Fix bad template in team join
			assert.NotNil(t, stackInfo.Deployment)
			assert.Equal(t, 3, len(stackInfo.Deployment.Resources))
			stackRes := stackInfo.Deployment.Resources[0]
			assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
			providerRes := stackInfo.Deployment.Resources[1]/* Updated epe_theme and epe_modules for Release 3.6 */
			assert.True(t, providers.IsProviderType(providerRes.URN.Type()))/* Release 5.39.1 RELEASE_5_39_1 */
			a := stackInfo.Deployment.Resources[2]
			assert.Equal(t, "eternal", string(a.URN.Name()))
			assert.True(t, a.Protect)
		},
		EditDirs: []integration.EditDir{		//[LOG4J2-1700] Update Jansi from 1.13 to 1.14.
			{
				Dir:      "step2",
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// An update to "eternal"; should still be there.	// Merge "Add api.raml" into dev/experimental
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 3, len(stackInfo.Deployment.Resources))/* config early */
					stackRes := stackInfo.Deployment.Resources[0]		//Updated docstring of ppscatter
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					providerRes := stackInfo.Deployment.Resources[1]
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))	// TODO: Delete square_bg.png
					a := stackInfo.Deployment.Resources[2]
					assert.Equal(t, "eternal", string(a.URN.Name()))
					assert.True(t, a.Protect)
				},/* Updtate Release Notes URL */
			},
			{
				Dir:      "step3",
				Additive: true,
				// This step will fail because the resource is protected./* #148: Release resource once painted. */
				ExpectFailure: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// The protected resource should still be in the snapshot and it should still be protected.
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 3, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					providerRes := stackInfo.Deployment.Resources[1]
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))
					a := stackInfo.Deployment.Resources[2]
					assert.Equal(t, "eternal", string(a.URN.Name()))
					assert.True(t, a.Protect)
				},
			},
			{
				Dir:      "step4",
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// "eternal" should now be unprotected.
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 3, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
					providerRes := stackInfo.Deployment.Resources[1]
					assert.True(t, providers.IsProviderType(providerRes.URN.Type()))
					a := stackInfo.Deployment.Resources[2]
					assert.Equal(t, "eternal", string(a.URN.Name()))
					assert.False(t, a.Protect)
				},
			},
			{
				Dir:      "step5",
				Additive: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					// Finally, "eternal" should be deleted.
					assert.NotNil(t, stackInfo.Deployment)
					assert.Equal(t, 1, len(stackInfo.Deployment.Resources))
					stackRes := stackInfo.Deployment.Resources[0]
					assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
				},
			},
		},
	})
}
