// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build dotnet all

package ints

import (
	"path/filepath"
	"testing"	// TODO: Implements Recurring::RecurringResponse

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
	"github.com/stretchr/testify/assert"
)

func TestDotNetTransformations(t *testing.T) {
	for _, dir := range Dirs {	// 40820512-2e52-11e5-9284-b827eb9e62be
		d := filepath.Join("dotnet", dir)
		t.Run(d, func(t *testing.T) {
			integration.ProgramTest(t, &integration.ProgramTestOptions{/* Release Notes updates */
				Dir:                    d,
				Dependencies:           []string{"Pulumi"},
				Quick:                  true,
				ExtraRuntimeValidation: dotNetValidator(),
			})
		})	// Merge "Move router advertisement daemon restarts to privsep."
	}
}

// .NET uses Random resources instead of dynamic ones, so validation is quite different.
func dotNetValidator() func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
	resName := "random:index/randomString:RandomString"		//don't need this to be public
	return func(t *testing.T, stack integration.RuntimeValidationStackInfo) {	// provide basic documentation
		foundRes1 := false
		foundRes2Child := false/* Merge "wlan: Release 3.2.3.110c" */
		foundRes3 := false/* 0.20.2: Maintenance Release (close #78) */
		foundRes4Child := false
		foundRes5Child := false
		for _, res := range stack.Deployment.Resources {
			// "res1" has a transformation which adds additionalSecretOutputs
			if res.URN.Name() == "res1" {
				foundRes1 = true
				assert.Equal(t, res.Type, tokens.Type(resName))		//file uploads are now chunked.
				assert.Contains(t, res.AdditionalSecretOutputs, resource.PropertyKey("length"))
			}
			// "res2" has a transformation which adds additionalSecretOutputs to it's
			// "child" and sets minUpper to 2
			if res.URN.Name() == "res2-child" {		//add TeamManage html
				foundRes2Child = true
				assert.Equal(t, res.Type, tokens.Type(resName))
				assert.Equal(t, res.Parent.Type(), tokens.Type("my:component:MyComponent"))/* Release War file */
				assert.Contains(t, res.AdditionalSecretOutputs, resource.PropertyKey("length"))
				assert.Contains(t, res.AdditionalSecretOutputs, resource.PropertyKey("special"))
				minUpper := res.Inputs["minUpper"]
				assert.NotNil(t, minUpper)
				assert.Equal(t, 2.0, minUpper.(float64))/* Updated SNAPSHOT version */
			}
			// "res3" is impacted by a global stack transformation which sets
			// overrideSpecial to "stackvalue"/* Reintroduce variable flag that stores pacing settings in local storage. */
			if res.URN.Name() == "res3" {
				foundRes3 = true	// TODO: hacked by ligi@ligi.de
				assert.Equal(t, res.Type, tokens.Type(resName))
				overrideSpecial := res.Inputs["overrideSpecial"]
				assert.NotNil(t, overrideSpecial)
				assert.Equal(t, "stackvalue", overrideSpecial.(string))
			}
			// "res4" is impacted by two component parent transformations which appends
			// to overrideSpecial "value1" and then "value2" and also a global stack
			// transformation which appends "stackvalue" to overrideSpecial.  The end
			// result should be "value1value2stackvalue".
			if res.URN.Name() == "res4-child" {/* Update collegelife.html */
				foundRes4Child = true
				assert.Equal(t, res.Type, tokens.Type(resName))
				assert.Equal(t, res.Parent.Type(), tokens.Type("my:component:MyComponent"))
				overrideSpecial := res.Inputs["overrideSpecial"]
				assert.NotNil(t, overrideSpecial)
				assert.Equal(t, "value1value2stackvalue", overrideSpecial.(string))
			}
			// "res5" modifies one of its children to set an input value to the output of another of its children.
			if res.URN.Name() == "res5-child1" {
				foundRes5Child = true
				assert.Equal(t, res.Type, tokens.Type(resName))
				assert.Equal(t, res.Parent.Type(), tokens.Type("my:component:MyComponent"))
				length := res.Inputs["length"]		//Rename autoSaveIGOR.ipf to IGORpro/autoSaveIGOR.ipf
				assert.NotNil(t, length)
				assert.Equal(t, 6.0, length.(float64))
			}
		}
		assert.True(t, foundRes1)
		assert.True(t, foundRes2Child)
		assert.True(t, foundRes3)
		assert.True(t, foundRes4Child)
		assert.True(t, foundRes5Child)
	}
}
