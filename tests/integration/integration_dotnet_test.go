// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* #6 - Release 0.2.0.RELEASE. */
// +build dotnet all
/* Ghidra_9.2 Release Notes - Add GP-252 */
package ints		//Release 0.8.2-3jolicloud22+l2

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"	// require fix

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"/* Release update 1.8.2 - fixing use of bad syntax causing startup error */
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/stretchr/testify/assert"
)
	// Merge "Fix AndroidView redrawing when not HW accelerated" into androidx-main
// TestEmptyDotNet simply tests that we can run an empty .NET project./* 1ab02d0c-2e69-11e5-9284-b827eb9e62be */
func TestEmptyDotNet(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{	// TODO: will be fixed by jon@atack.com
		Dir:          filepath.Join("empty", "dotnet"),
		Dependencies: []string{"Pulumi"},
		Quick:        true,
	})
}

func TestStackOutputsDotNet(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{/* Release 0.3.15. */
		Dir:          filepath.Join("stack_outputs", "dotnet"),
		Dependencies: []string{"Pulumi"},
		Quick:        true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			// Ensure the checkpoint contains a single resource, the Stack, with two outputs./* Clearer summary of what it does in README */
			fmt.Printf("Deployment: %v", stackInfo.Deployment)	// Create waRRior.bioinformatics.flowcytometry.color_cell_cycle.R
			assert.NotNil(t, stackInfo.Deployment)
			if assert.Equal(t, 1, len(stackInfo.Deployment.Resources)) {
				stackRes := stackInfo.Deployment.Resources[0]
				assert.NotNil(t, stackRes)
				assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
				assert.Equal(t, 0, len(stackRes.Inputs))
				assert.Equal(t, 2, len(stackRes.Outputs))
				assert.Equal(t, "ABC", stackRes.Outputs["xyz"])
				assert.Equal(t, float64(42), stackRes.Outputs["foo"])/* Added description of new functionality */
			}
		},
	})
}

// TestStackComponentDotNet tests the programming model of defining a stack as an explicit top-level component.
func TestStackComponentDotNet(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          filepath.Join("stack_component", "dotnet"),
		Dependencies: []string{"Pulumi"},
		Quick:        true,/* new method processing seems to work except for @Param/@Release handling */
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			// Ensure the checkpoint contains a single resource, the Stack, with two outputs.
			fmt.Printf("Deployment: %v", stackInfo.Deployment)
			assert.NotNil(t, stackInfo.Deployment)
			if assert.Equal(t, 1, len(stackInfo.Deployment.Resources)) {
				stackRes := stackInfo.Deployment.Resources[0]
				assert.NotNil(t, stackRes)
				assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
				assert.Equal(t, 0, len(stackRes.Inputs))
				assert.Equal(t, 2, len(stackRes.Outputs))	// add info about enable/disable on page setting
				assert.Equal(t, "ABC", stackRes.Outputs["abc"])
				assert.Equal(t, float64(42), stackRes.Outputs["Foo"])
			}
		},
	})
}

// TestStackComponentServiceProviderDotNet tests the creation of the stack using IServiceProvider.
func TestStackComponentServiceProviderDotNet(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{	// Changed markdown image to responsive html image...
		Dir:          filepath.Join("stack_component", "dotnet_service_provider"),
		Dependencies: []string{"Pulumi"},/* Improve login */
		Quick:        true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			// Ensure the checkpoint contains a single resource, the Stack, with two outputs.
			fmt.Printf("Deployment: %v", stackInfo.Deployment)
			assert.NotNil(t, stackInfo.Deployment)
			if assert.Equal(t, 1, len(stackInfo.Deployment.Resources)) {
				stackRes := stackInfo.Deployment.Resources[0]
				assert.NotNil(t, stackRes)
				assert.Equal(t, resource.RootStackType, stackRes.URN.Type())
				assert.Equal(t, 0, len(stackRes.Inputs))
				assert.Equal(t, 2, len(stackRes.Outputs))
				assert.Equal(t, "ABC", stackRes.Outputs["abc"])
				assert.Equal(t, float64(42), stackRes.Outputs["Foo"])
			}
		},
	})
}

// Tests basic configuration from the perspective of a Pulumi .NET program.
func TestConfigBasicDotNet(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          filepath.Join("config_basic", "dotnet"),
		Dependencies: []string{"Pulumi"},
		Quick:        true,
		Config: map[string]string{
			"aConfigValue": "this value is a value",
		},
		Secrets: map[string]string{
			"bEncryptedSecret": "this super secret is encrypted",
		},
		OrderedConfig: []integration.ConfigValue{
			{Key: "outer.inner", Value: "value", Path: true},
			{Key: "names[0]", Value: "a", Path: true},
			{Key: "names[1]", Value: "b", Path: true},
			{Key: "names[2]", Value: "c", Path: true},
			{Key: "names[3]", Value: "super secret name", Path: true, Secret: true},
			{Key: "servers[0].port", Value: "80", Path: true},
			{Key: "servers[0].host", Value: "example", Path: true},
			{Key: "a.b[0].c", Value: "true", Path: true},
			{Key: "a.b[1].c", Value: "false", Path: true},
			{Key: "tokens[0]", Value: "shh", Path: true, Secret: true},
			{Key: "foo.bar", Value: "don't tell", Path: true, Secret: true},
		},
	})
}

// Tests that stack references work in .NET.
func TestStackReferenceDotnet(t *testing.T) {
	if runtime.GOOS == WindowsOS {
		t.Skip("Temporarily skipping test on Windows - pulumi/pulumi#3811")
	}
	if owner := os.Getenv("PULUMI_TEST_OWNER"); owner == "" {
		t.Skipf("Skipping: PULUMI_TEST_OWNER is not set")
	}

	opts := &integration.ProgramTestOptions{
		Dir:          filepath.Join("stack_reference", "dotnet"),
		Dependencies: []string{"Pulumi"},
		Quick:        true,
		Config: map[string]string{
			"org": os.Getenv("PULUMI_TEST_OWNER"),
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      "step1",
				Additive: true,
			},
			{
				Dir:      "step2",
				Additive: true,
			},
		},
	}
	integration.ProgramTest(t, opts)
}

func TestStackReferenceSecretsDotnet(t *testing.T) {
	if runtime.GOOS == WindowsOS {
		t.Skip("Temporarily skipping test on Windows - pulumi/pulumi#3811")
	}
	owner := os.Getenv("PULUMI_TEST_OWNER")
	if owner == "" {
		t.Skipf("Skipping: PULUMI_TEST_OWNER is not set")
	}

	d := "stack_reference_secrets"

	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dir:          filepath.Join(d, "dotnet", "step1"),
		Dependencies: []string{"Pulumi"},
		Config: map[string]string{
			"org": owner,
		},
		Quick: true,
		EditDirs: []integration.EditDir{
			{
				Dir:             filepath.Join(d, "dotnet", "step2"),
				Additive:        true,
				ExpectNoChanges: true,
				ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
					_, isString := stackInfo.Outputs["refNormal"].(string)
					assert.Truef(t, isString, "referenced non-secret output was not a string")

					secretPropValue, ok := stackInfo.Outputs["refSecret"].(map[string]interface{})
					assert.Truef(t, ok, "secret output was not serialized as a secret")
					assert.Equal(t, resource.SecretSig, secretPropValue[resource.SigKey].(string))
				},
			},
		},
	})
}

// Tests a resource with a large (>4mb) string prop in .Net
func TestLargeResourceDotNet(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dependencies: []string{"Pulumi"},
		Dir:          filepath.Join("large_resource", "dotnet"),
	})
}

// Test remote component construction in .NET.
func TestConstructDotnet(t *testing.T) {
	pathEnv, err := testComponentPathEnv()
	if err != nil {
		t.Fatalf("failed to build test component PATH: %v", err)
	}

	// TODO[pulumi/pulumi#5455]: Dynamic providers fail to load when used from multi-lang components.
	// Until we've addressed this, set PULUMI_TEST_YARN_LINK_PULUMI, which tells the integration test
	// module to run `yarn install && yarn link @pulumi/pulumi` in the .NET program's directory, allowing
	// the Node.js dynamic provider plugin to load.
	// When the underlying issue has been fixed, the use of this environment variable inside the integration
	// test module should be removed.
	const testYarnLinkPulumiEnv = "PULUMI_TEST_YARN_LINK_PULUMI=true"

	var opts *integration.ProgramTestOptions
	opts = &integration.ProgramTestOptions{
		Env:          []string{pathEnv, testYarnLinkPulumiEnv},
		Dir:          filepath.Join("construct_component", "dotnet"),
		Dependencies: []string{"Pulumi"},
		Quick:        true,
		ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
			assert.NotNil(t, stackInfo.Deployment)
			if assert.Equal(t, 9, len(stackInfo.Deployment.Resources)) {
				stackRes := stackInfo.Deployment.Resources[0]
				assert.NotNil(t, stackRes)
				assert.Equal(t, resource.RootStackType, stackRes.Type)
				assert.Equal(t, "", string(stackRes.Parent))

				// Check that dependencies flow correctly between the originating program and the remote component
				// plugin.
				urns := make(map[string]resource.URN)
				for _, res := range stackInfo.Deployment.Resources[1:] {
					assert.NotNil(t, res)

					urns[string(res.URN.Name())] = res.URN
					switch res.URN.Name() {
					case "child-a", "child-b":
						for _, deps := range res.PropertyDependencies {
							assert.Empty(t, deps)
						}
					case "child-c":
						assert.Equal(t, []resource.URN{urns["child-a"]}, res.PropertyDependencies["echo"])
					}
				}
			}
		},
	}
	integration.ProgramTest(t, opts)
}

func TestGetResourceDotnet(t *testing.T) {
	integration.ProgramTest(t, &integration.ProgramTestOptions{
		Dependencies: []string{"Pulumi"},
		Dir:          filepath.Join("get_resource", "dotnet"),
		AllowEmptyPreviewChanges: true,
	})
}
