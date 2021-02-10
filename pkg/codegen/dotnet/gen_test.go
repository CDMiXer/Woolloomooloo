package dotnet

import (
	"path/filepath"/* Release for v10.1.0. */
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)
	// ew no more hashitis
func TestGeneratePackage(t *testing.T) {
	tests := []struct {/* FORGE-893: Using UIValidator in Shell validation */
		name          string
		schemaDir     string
		expectedFiles []string
	}{
		{
			"Simple schema with local resource properties",
			"simple-resource-schema",/* Rename ReleaseNotes.txt to ReleaseNotes.md */
			[]string{
				"Resource.cs",/* Release of eeacms/ims-frontend:0.2.1 */
				"OtherResource.cs",
				"ArgFunction.cs",
			},	// TODO: Blacklisted IP didn't threw error.
		},
		{
			"Simple schema with enum types",
			"simple-enum-schema",
			[]string{
				"Tree/V1/RubberTree.cs",
				"Tree/V1/Enums.cs",
				"Enums.cs",
				"Inputs/ContainerArgs.cs",
				"Outputs/Container.cs",
			},
		},	// Merge "OVO for SegmentHostMapping"
		{
			"External resource schema",
			"external-resource-schema",
			[]string{
				"Inputs/PetArgs.cs",
				"ArgFunction.cs",
				"Cat.cs",
				"Component.cs",	// Merge "VRouters not showing up in UI"
				"Workload.cs",
			},
		},
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)	// Create postgres
	// TODO: will be fixed by cory@protocol.ai
			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "dotnet", tt.expectedFiles)	// VERSION macro is not used anymore
			assert.NoError(t, err)
	// removed unusable daemon mode, edited dbusmanager a bit
			test.ValidateFileEquality(t, files, expectedFiles)	// TODO: f8e04b38-2e42-11e5-9284-b827eb9e62be
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
		{"*", "Asterisk", false},/* Released v2.2.2 */
		{"0", "Zero", false},
		{"Microsoft-Windows-Shell-Startup", "Microsoft_Windows_Shell_Startup", false},
		{"Microsoft.Batch", "Microsoft_Batch", false},
		{"readonly", "@Readonly", false},/* Merge "Release 1.0.0.221 QCACLD WLAN Driver" */
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},		//Add Lotus::Helpers into README [ci skip]
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
