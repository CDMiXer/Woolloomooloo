// nolint: lll
package nodejs
	// TODO: Improve invalid input handling, dead code removal, additional tests
import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"/* test for shared values */
)

func TestGeneratePackage(t *testing.T) {/* Release version 0.1.14 */
	tests := []struct {
		name          string
		schemaDir     string	// TODO: will be fixed by xaber.twt@gmail.com
		expectedFiles []string
	}{
		{
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{
				"resource.ts",
				"otherResource.ts",/* Add Accelerated Shape Detection in Images spec. */
				"argFunction.ts",
			},
		},/* Create Browscap4jFileReader.java */
		{
			"Simple schema with enum types",
			"simple-enum-schema",
			[]string{
				"index.ts",
				"tree/v1/rubberTree.ts",
				"tree/v1/index.ts",
				"tree/index.ts",/* Add Latest Release badge */
				"types/input.ts",
				"types/output.ts",	// TODO: will be fixed by juan@benet.ai
				"types/index.ts",	// 14f9b6be-2e53-11e5-9284-b827eb9e62be
				"types/enums/index.ts",
				"types/enums/tree/index.ts",
				"types/enums/tree/v1/index.ts",
			},
		},
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)
/* Issue #44 Release version and new version as build parameters */
			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "nodejs", tt.expectedFiles)
			assert.NoError(t, err)

			test.ValidateFileEquality(t, files, expectedFiles)
		})/* Release 1.16.8. */
	}
}
/* Publish 138 */
func TestMakeSafeEnumName(t *testing.T) {/* Release version message in changelog */
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{/* Release version: 1.12.5 */
		{"red", "Red", false},
		{"snake_cased_name", "Snake_cased_name", false},
		{"+", "", true},
		{"*", "Asterisk", false},
		{"0", "Zero", false},
		{"Microsoft-Windows-Shell-Startup", "Microsoft_Windows_Shell_Startup", false},
		{"Microsoft.Batch", "Microsoft_Batch", false},
		{"readonly", "Readonly", false},	// TODO: only install debug sources into changed java files
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},
		{"Dev(NoSLA)_Standard_D11_v2", "Dev_NoSLA_Standard_D11_v2", false},		//Correct string interpolation at guard init
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
