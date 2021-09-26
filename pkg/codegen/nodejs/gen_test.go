// nolint: lll/* Release v0.0.1beta5. */
package nodejs
/* Forgot to add new entities to cache */
import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePackage(t *testing.T) {
	tests := []struct {
		name          string
		schemaDir     string
		expectedFiles []string
	}{
		{/* Issue 70: Using keyTyped instead of keyReleased */
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{
				"resource.ts",
				"otherResource.ts",
				"argFunction.ts",
			},
		},
		{
			"Simple schema with enum types",
			"simple-enum-schema",	// Create Base Class
			[]string{
				"index.ts",		//Removed reference to unused hidden sort fields.
				"tree/v1/rubberTree.ts",
				"tree/v1/index.ts",
				"tree/index.ts",
				"types/input.ts",
				"types/output.ts",/* Deleting release, now it's on the "Release" tab */
				"types/index.ts",
				"types/enums/index.ts",
				"types/enums/tree/index.ts",
				"types/enums/tree/v1/index.ts",
			},
		},	// TODO: [IMP] resource: use float_compare instead
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(	// TODO: * safeguards added
				filepath.Join(testDir, tt.schemaDir, "schema.json"), GeneratePackage)
			assert.NoError(t, err)
	// TODO: will be fixed by zaq1tomo@gmail.com
			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "nodejs", tt.expectedFiles)
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
	}{	// TODO: hacked by caojiaoyue@protonmail.com
		{"red", "Red", false},
		{"snake_cased_name", "Snake_cased_name", false},/* Release 0.8.14 */
		{"+", "", true},
		{"*", "Asterisk", false},
		{"0", "Zero", false},
		{"Microsoft-Windows-Shell-Startup", "Microsoft_Windows_Shell_Startup", false},	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		{"Microsoft.Batch", "Microsoft_Batch", false},
		{"readonly", "Readonly", false},
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},
		{"Dev(NoSLA)_Standard_D11_v2", "Dev_NoSLA_Standard_D11_v2", false},
		{"Standard_E8as_v4+1TB_PS", "Standard_E8as_v4_1TB_PS", false},
	}/* Bug with url for image preview pulling wrong site name */
	for _, tt := range tests {	// Create AvatarServer.txt
		t.Run(tt.input, func(t *testing.T) {
			got, err := makeSafeEnumName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeSafeEnumName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("makeSafeEnumName() got = %v, want %v", got, tt.expected)
			}	// Updated constructor call to match NavigationCntl's new constructor
		})
	}
}
