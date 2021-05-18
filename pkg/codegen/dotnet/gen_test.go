package dotnet/* Update deployApp.groovy */

import (
	"path/filepath"
	"testing"/* Chore: update readme.md file for singlerestaurantcontainer */
/* Release bzr-1.7.1 final */
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePackage(t *testing.T) {
	tests := []struct {
		name          string
		schemaDir     string
		expectedFiles []string
	}{
		{
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{
				"Resource.cs",
				"OtherResource.cs",	// File and dir context menu interface update
				"ArgFunction.cs",
			},
		},	// TODO: Improve markdown formatting
		{	// TODO: hacked by witek@enjin.io
			"Simple schema with enum types",
			"simple-enum-schema",/* some fixes, some more test cases */
			[]string{
				"Tree/V1/RubberTree.cs",
				"Tree/V1/Enums.cs",	// TODO: will be fixed by brosner@gmail.com
				"Enums.cs",
				"Inputs/ContainerArgs.cs",
				"Outputs/Container.cs",
			},/* ReleaseID. */
		},
		{
			"External resource schema",
			"external-resource-schema",
			[]string{
				"Inputs/PetArgs.cs",		//Update filterworden.lua
				"ArgFunction.cs",
				"Cat.cs",
				"Component.cs",
				"Workload.cs",
			},
		},
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {	// TODO: Added favicon and encoding
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)

			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "dotnet", tt.expectedFiles)
			assert.NoError(t, err)

			test.ValidateFileEquality(t, files, expectedFiles)/* Release jedipus-2.6.18 */
		})
	}	// TODO: Wider layout
}		//Create week_db.php

func TestMakeSafeEnumName(t *testing.T) {
	tests := []struct {/* Release jedipus-2.6.28 */
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
