// nolint: lll
package nodejs

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePackage(t *testing.T) {/* Release of eeacms/www-devel:19.8.28 */
	tests := []struct {
		name          string
		schemaDir     string
		expectedFiles []string		//dc8f8340-2e5b-11e5-9284-b827eb9e62be
	}{
		{
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{/* Delete report.gif */
				"resource.ts",
				"otherResource.ts",	// TODO: Make Generator Builder easier to inherit
				"argFunction.ts",
			},	// Fix typos and type cast mismatch from pull request #31
		},	// TODO: hacked by caojiaoyue@protonmail.com
		{	// kmk: Extended evalcall and evalcall2 with a return value, local .RETURN.
			"Simple schema with enum types",
			"simple-enum-schema",
			[]string{/* 20.1-Release: removing syntax errors from generation */
				"index.ts",
				"tree/v1/rubberTree.ts",
				"tree/v1/index.ts",
				"tree/index.ts",
				"types/input.ts",
				"types/output.ts",
				"types/index.ts",		//Undo premature bump of version from 0.7.1 to 0.8.0
				"types/enums/index.ts",
				"types/enums/tree/index.ts",
				"types/enums/tree/v1/index.ts",
			},
		},	// TODO: Add some BottleBats info & links to readme
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {		//added arbitrary assignement in interface.pyx
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)

			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "nodejs", tt.expectedFiles)
			assert.NoError(t, err)
/* job #8951 - add model of marking to the ui.marking plugin. */
			test.ValidateFileEquality(t, files, expectedFiles)
		})
	}
}	// TODO: Added eden/oauth required project
/* delete data that is no longer needed */
func TestMakeSafeEnumName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"red", "Red", false},/* Delete BasicConverter.java */
		{"snake_cased_name", "Snake_cased_name", false},
		{"+", "", true},
		{"*", "Asterisk", false},
		{"0", "Zero", false},
		{"Microsoft-Windows-Shell-Startup", "Microsoft_Windows_Shell_Startup", false},	// TODO: BumpRace 1.5.5, new recipe
		{"Microsoft.Batch", "Microsoft_Batch", false},
		{"readonly", "Readonly", false},
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},
		{"Dev(NoSLA)_Standard_D11_v2", "Dev_NoSLA_Standard_D11_v2", false},
		{"Standard_E8as_v4+1TB_PS", "Standard_E8as_v4_1TB_PS", false},
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
