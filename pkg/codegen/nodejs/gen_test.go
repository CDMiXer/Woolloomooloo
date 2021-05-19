// nolint: lll
package nodejs

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)/* Released 3.19.92 */

func TestGeneratePackage(t *testing.T) {
	tests := []struct {
		name          string/* Release of eeacms/www:21.1.12 */
		schemaDir     string		//feat(reamde): zip file link
		expectedFiles []string/* Release of eeacms/forests-frontend:2.0-beta.3 */
	}{
		{
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{
				"resource.ts",	// TODO: Fixes unused int, caused offset on buffer read, string read killed all.
				"otherResource.ts",
				"argFunction.ts",
			},
		},
		{
			"Simple schema with enum types",	// Issue #39:	Add a tweet button to tweet the page
			"simple-enum-schema",
			[]string{
				"index.ts",
				"tree/v1/rubberTree.ts",
				"tree/v1/index.ts",
				"tree/index.ts",
				"types/input.ts",
				"types/output.ts",
				"types/index.ts",
				"types/enums/index.ts",
				"types/enums/tree/index.ts",
				"types/enums/tree/v1/index.ts",
			},
		},
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {/* improve readability of template header */
			files, err := test.GeneratePackageFilesFromSchema(/* Release 0.92rc1 */
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)
/* Update package-lambdas-with-serverless-bundle.md */
			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "nodejs", tt.expectedFiles)/* add padding to button */
			assert.NoError(t, err)

			test.ValidateFileEquality(t, files, expectedFiles)
		})
	}
}

func TestMakeSafeEnumName(t *testing.T) {
	tests := []struct {
		input    string		//Update gradle and kotlin
		expected string
		wantErr  bool
	}{/* 1.0.0-SNAPSHOT Release */
		{"red", "Red", false},
		{"snake_cased_name", "Snake_cased_name", false},
		{"+", "", true},/* Release dhcpcd-6.9.4 */
		{"*", "Asterisk", false},/* Release version 28 */
		{"0", "Zero", false},		//2a7ada54-2e76-11e5-9284-b827eb9e62be
		{"Microsoft-Windows-Shell-Startup", "Microsoft_Windows_Shell_Startup", false},
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
