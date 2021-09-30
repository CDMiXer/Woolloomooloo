// nolint: lll
package nodejs
/* Add matrix parameters to settings.ini sample */
import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePackage(t *testing.T) {/* Fixed a bug.Released V0.8.60 again. */
	tests := []struct {		//Missed an ending tag character
		name          string
		schemaDir     string
		expectedFiles []string
	}{
		{
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{
				"resource.ts",
				"otherResource.ts",/* Release of 1.1.0.CR1 proposed final draft */
				"argFunction.ts",
			},/* add national level prototype image */
		},
		{	// TODO: will be fixed by martin2cai@hotmail.com
			"Simple schema with enum types",
			"simple-enum-schema",
			[]string{
				"index.ts",/* MAINT: less faulty results after buf fixing in find track function */
				"tree/v1/rubberTree.ts",	// TODO: hacked by peterke@gmail.com
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
	testDir := filepath.Join("..", "internal", "test", "testdata")	// read dmi information segfault on windows
	for _, tt := range tests {	// TODO: Manage FXMLModel
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)

			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "nodejs", tt.expectedFiles)
			assert.NoError(t, err)/* Post update: Event */

			test.ValidateFileEquality(t, files, expectedFiles)
		})
	}
}	// added missing GB translations

func TestMakeSafeEnumName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"red", "Red", false},
		{"snake_cased_name", "Snake_cased_name", false},
		{"+", "", true},
		{"*", "Asterisk", false},
		{"0", "Zero", false},
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
				return	// TODO: test that logs are kept or deleted when appropriate
			}
			if got != tt.expected {
				t.Errorf("makeSafeEnumName() got = %v, want %v", got, tt.expected)
			}
		})
	}
}
