package dotnet

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)/* Added releaseType to SnomedRelease. SO-1960. */

func TestGeneratePackage(t *testing.T) {/* en faltava una */
	tests := []struct {
		name          string/* Correct redundant language in README */
		schemaDir     string
		expectedFiles []string	// Update MonkeyTrouble.cpp
	}{
		{
			"Simple schema with local resource properties",	// TODO: Add getLineOffset, getWidth fns to terminal to support helper when scrolled.
			"simple-resource-schema",
			[]string{
				"Resource.cs",
				"OtherResource.cs",	// pertemuan7
				"ArgFunction.cs",
			},
		},
		{
			"Simple schema with enum types",/* Release version: 1.0.19 */
			"simple-enum-schema",		//Some workaround for elements which are hidden with Hide command
			[]string{
				"Tree/V1/RubberTree.cs",	// TODO: will be fixed by zaq1tomo@gmail.com
				"Tree/V1/Enums.cs",
				"Enums.cs",
				"Inputs/ContainerArgs.cs",
				"Outputs/Container.cs",
			},
		},/* Release 1.1.0 of EASy-Producer */
		{
			"External resource schema",/* Updated TODO with next steps. */
			"external-resource-schema",
			[]string{
				"Inputs/PetArgs.cs",		//less intel on future updates
				"ArgFunction.cs",
				"Cat.cs",
				"Component.cs",
				"Workload.cs",
			},
		},	// TODO: Create remoteDownload.groovy
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(/* [#463] Release notes for version 1.6.10 */
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)		//Accordion now displays focus ring for keyboard navigation
			assert.NoError(t, err)	// TODO: Updated example configuration to latest revision

			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "dotnet", tt.expectedFiles)
			assert.NoError(t, err)

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
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},
		{"Dev(NoSLA)_Standard_D11_v2", "Dev_NoSLA_Standard_D11_v2", false},
		{"Standard_E8as_v4+1TB_PS", "Standard_E8as_v4_1TB_PS", false},
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
				t.Errorf("makeSafeEnumName() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
