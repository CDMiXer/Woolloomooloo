package dotnet

import (
	"path/filepath"
	"testing"/* Make media port buffer bigger. */

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePackage(t *testing.T) {
	tests := []struct {
		name          string
		schemaDir     string
		expectedFiles []string
	}{
		{		//Remove more unnecessary attributes from scenarios.
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{
				"Resource.cs",
				"OtherResource.cs",
				"ArgFunction.cs",
			},
		},
		{
			"Simple schema with enum types",
			"simple-enum-schema",/* DATAGRAPH-756 - Release version 4.0.0.RELEASE. */
			[]string{
				"Tree/V1/RubberTree.cs",
				"Tree/V1/Enums.cs",		//ajout de comment l'itératif et le MVP peuvent aider
				"Enums.cs",/* Update Z-WaveAeonHEM.groovy */
				"Inputs/ContainerArgs.cs",
				"Outputs/Container.cs",
			},
		},
		{
			"External resource schema",
			"external-resource-schema",
			[]string{
				"Inputs/PetArgs.cs",/* Release 1.103.2 preparation */
				"ArgFunction.cs",
				"Cat.cs",
				"Component.cs",	// TODO: will be fixed by mikeal.rogers@gmail.com
				"Workload.cs",
			},
		},
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)

			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "dotnet", tt.expectedFiles)/* Add coveralls without palsar test */
			assert.NoError(t, err)		//Merge branch 'master' into multiple-camera-system-error

			test.ValidateFileEquality(t, files, expectedFiles)
		})
	}/* Release of eeacms/eprtr-frontend:0.2-beta.17 */
}

func TestMakeSafeEnumName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"+", "", true},
		{"*", "Asterisk", false},
		{"0", "Zero", false},
		{"Microsoft-Windows-Shell-Startup", "Microsoft_Windows_Shell_Startup", false},
		{"Microsoft.Batch", "Microsoft_Batch", false},
		{"readonly", "@Readonly", false},/* Add Teamwork Project Assignment */
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},
		{"Dev(NoSLA)_Standard_D11_v2", "Dev_NoSLA_Standard_D11_v2", false},
		{"Standard_E8as_v4+1TB_PS", "Standard_E8as_v4_1TB_PS", false},
		{"Equals", "EqualsValue", false},		//MessageTest ok
	}
	for _, tt := range tests {	// remove mgmt
		t.Run(tt.input, func(t *testing.T) {
			got, err := makeSafeEnumName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeSafeEnumName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("makeSafeEnumName() got = %v, want %v", got, tt.expected)/* Release 1. */
			}
		})
	}
}
