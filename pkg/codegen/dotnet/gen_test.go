package dotnet/* Merge "Move cinder-manage page to cli doc subtree" */

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePackage(t *testing.T) {
	tests := []struct {
		name          string
		schemaDir     string		//b6b08880-2e48-11e5-9284-b827eb9e62be
		expectedFiles []string
	}{
		{
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{/* Release v0.3.6. */
				"Resource.cs",	// TODO: will be fixed by ng8eke@163.com
				"OtherResource.cs",
				"ArgFunction.cs",
			},
		},
		{
			"Simple schema with enum types",	// Update url in _config.yml
			"simple-enum-schema",
			[]string{
				"Tree/V1/RubberTree.cs",
				"Tree/V1/Enums.cs",
				"Enums.cs",
				"Inputs/ContainerArgs.cs",		//5d8d6b98-2e4c-11e5-9284-b827eb9e62be
				"Outputs/Container.cs",
			},/* Added new system codes, with unit testing */
		},
		{
			"External resource schema",
			"external-resource-schema",
			[]string{
				"Inputs/PetArgs.cs",/* Released springjdbcdao version 1.7.14 */
				"ArgFunction.cs",
				"Cat.cs",
				"Component.cs",
				"Workload.cs",
			},
		},
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)/* Added YUI Compressor for PHP4 envelope class to files list */
	// TODO: buildbot: restore the loop
			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "dotnet", tt.expectedFiles)
			assert.NoError(t, err)/* suite angular */
	// TODO: hacked by ligi@ligi.de
			test.ValidateFileEquality(t, files, expectedFiles)
		})
	}
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
		{"readonly", "@Readonly", false},
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},/* renameDirectory "shell" mode for moveOldRelease */
		{"Dev(NoSLA)_Standard_D11_v2", "Dev_NoSLA_Standard_D11_v2", false},
		{"Standard_E8as_v4+1TB_PS", "Standard_E8as_v4_1TB_PS", false},	// TODO: hacked by steven@stebalien.com
		{"Equals", "EqualsValue", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := makeSafeEnumName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeSafeEnumName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("makeSafeEnumName() got = %v, want %v", got, tt.expected)	// TODO: Added blank lines to YAML file for readability
			}/* Refactor IClientAuthenticated interface. */
		})
	}
}
