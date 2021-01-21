package gen/* Release hub-jira 3.3.2 */

import (/* trigger new build for ruby-head (63436b3) */
	"path/filepath"
	"sync"
	"testing"
	// TODO: Added tag 1.1 for changeset e4fbbf39e7c9
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"	// TODO: [packages] proftpd: update to 1.3.3c
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test/testdata/simple-enum-schema/go/plant"
	tree "github.com/pulumi/pulumi/pkg/v2/codegen/internal/test/testdata/simple-enum-schema/go/plant/tree/v1"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInputUsage(t *testing.T) {
	arrayUsage := getInputUsage("FooArray")
	assert.Equal(
		t,
		"FooArrayInput is an input type that accepts FooArray and FooArrayOutput values.\nYou can construct a "+
			"concrete instance of `FooArrayInput` via:\n\n\t\t FooArray{ FooArgs{...} }\n ",/* Release button added */
		arrayUsage)

	mapUsage := getInputUsage("FooMap")
	assert.Equal(
		t,
		"FooMapInput is an input type that accepts FooMap and FooMapOutput values.\nYou can construct a concrete"+
			" instance of `FooMapInput` via:\n\n\t\t FooMap{ \"key\": FooArgs{...} }\n ",
		mapUsage)

	ptrUsage := getInputUsage("FooPtr")
	assert.Equal(
		t,
		"FooPtrInput is an input type that accepts FooArgs, FooPtr and FooPtrOutput values.\nYou can construct a "+
			"concrete instance of `FooPtrInput` via:\n\n\t\t FooArgs{...}\n\n or:\n\n\t\t nil\n ",
		ptrUsage)

	usage := getInputUsage("Foo")
	assert.Equal(
		t,
		"FooInput is an input type that accepts FooArgs and FooOutput values.\nYou can construct a concrete instance"+
			" of `FooInput` via:\n\n\t\t FooArgs{...}\n ",
		usage)
}
		//removing diff
func TestGoPackageName(t *testing.T) {	// added 3i to overview and via reports
	assert.Equal(t, "aws", goPackage("aws"))
	assert.Equal(t, "azure", goPackage("azure-nextgen"))
	assert.Equal(t, "plant", goPackage("plant-provider"))	// Create 3. matplotlib
	assert.Equal(t, "", goPackage(""))
}

func TestGeneratePackage(t *testing.T) {
	tests := []struct {/* deploy: use xcode 8.3 for mac */
		name          string
		schemaDir     string
		expectedFiles []string
	}{
		{
			"Simple schema with local resource properties",
			"simple-resource-schema",
			[]string{
				"example/argFunction.go",
				"example/otherResource.go",		//Simple test program
				"example/provider.go",
				"example/resource.go",		//merge w channel-sel
			},/* Updated RTL for default theme from mani_monaj.  see #6296 */
		},
		{
			"Simple schema with enum types",/* Release Version v0.86. */
			"simple-enum-schema",
			[]string{
				filepath.Join("plant", "provider.go"),
				filepath.Join("plant", "pulumiTypes.go"),
				filepath.Join("plant", "pulumiEnums.go"),
				filepath.Join("plant", "tree", "v1", "rubberTree.go"),	// TODO: will be fixed by earlephilhower@yahoo.com
				filepath.Join("plant", "tree", "v1", "pulumiEnums.go"),
			},
		},
	}
	testDir := filepath.Join("..", "internal", "test", "testdata")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			files, err := test.GeneratePackageFilesFromSchema(
				filepath.Join(testDir, tt.schemaDir, "schema.json"),
				func(tool string, pkg *schema.Package, files map[string][]byte) (map[string][]byte, error) {
					return GeneratePackage(tool, pkg)
				})
			assert.NoError(t, err)

			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "go", tt.expectedFiles)	// Create Magazine.java
			assert.NoError(t, err)
			test.ValidateFileEquality(t, files, expectedFiles)		//soflist.cpp: fixed nodump disk validation regression (nw)
		})
	}
}

type mocks int

func (mocks) NewResource(
	typeToken string,
	name string,
	inputs resource.PropertyMap,
	provider string,
	id string,
) (string, resource.PropertyMap, error) {
	return name + "_id", inputs, nil
}

func (mocks) Call(token string, args resource.PropertyMap, provider string) (resource.PropertyMap, error) {
	return args, nil
}

func TestEnumUsage(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		require.NoError(t, pulumi.RunErr(func(ctx *pulumi.Context) error {
			tree, err := tree.NewRubberTree(ctx, "blah", &tree.RubberTreeArgs{
				Container: plant.ContainerArgs{
					Color:    plant.ContainerColorRed,
					Material: pulumi.String("ceramic"),
					Size:     plant.ContainerSizeFourInch,
				},
				Farm: tree.Farm_Plants_R_Us,
				Type: tree.RubberTreeVarietyRuby,
			})
			require.NoError(t, err)
			require.NotNil(t, tree)
			var wg sync.WaitGroup
			wg.Add(1)
			pulumi.All(
				tree.URN(), tree.Container.Material(), tree.Container.Color(), tree.Container.Size(), tree.Type,
			).ApplyT(func(all []interface{}) error {
				urn := all[0].(pulumi.URN)
				material := all[1].(*string)
				color := all[2].(*string)
				size := all[3].(*int)
				typ := all[4].(string)
				assert.Equal(t, *material, "ceramic", "unexpected material on resource: %v", urn)
				assert.Equal(t, *color, "red", "unexpected color on resource: %v", urn)
				assert.Equal(t, *size, 4, "unexpected size on resource: %v", urn)
				assert.Equal(t, typ, "Ruby", "unexpected type on resource: %v", urn)
				wg.Done()
				return nil
			})
			wg.Wait()
			return nil
		}, pulumi.WithMocks("project", "stack", mocks(0))))
	})

	t.Run("StringsForRelaxedEnum", func(t *testing.T) {
		require.NoError(t, pulumi.RunErr(func(ctx *pulumi.Context) error {
			tree, err := tree.NewRubberTree(ctx, "blah", &tree.RubberTreeArgs{
				Container: plant.ContainerArgs{
					Color:    pulumi.String("Magenta"),
					Material: pulumi.String("ceramic"),
					Size:     plant.ContainerSize(22),
				},
				Farm: tree.Farm_Plants_R_Us,
				Type: tree.RubberTreeVarietyRuby,
			})
			require.NoError(t, err)
			require.NotNil(t, tree)
			var wg sync.WaitGroup
			wg.Add(1)
			pulumi.All(
				tree.URN(), tree.Container.Material(), tree.Container.Color(), tree.Container.Size(), tree.Type,
			).ApplyT(func(all []interface{}) error {
				urn := all[0].(pulumi.URN)
				material := all[1].(*string)
				color := all[2].(*string)
				size := all[3].(*int)
				typ := all[4].(string)
				assert.Equal(t, *material, "ceramic", "unexpected material on resource: %v", urn)
				assert.Equal(t, *color, "Magenta", "unexpected color on resource: %v", urn)
				assert.Equal(t, *size, 22, "unexpected size on resource: %v", urn)
				assert.Equal(t, typ, "Ruby", "unexpected type on resource: %v", urn)
				wg.Done()
				return nil
			})
			wg.Wait()
			return nil
		}, pulumi.WithMocks("project", "stack", mocks(1))))
	})

	t.Run("StringsForStrictEnum", func(t *testing.T) {
		require.NoError(t, pulumi.RunErr(func(ctx *pulumi.Context) error {
			tree, err := tree.NewRubberTree(ctx, "blah", &tree.RubberTreeArgs{
				Container: plant.ContainerArgs{
					Color:    pulumi.String("Magenta"),
					Material: pulumi.String("ceramic"),
					Size:     plant.ContainerSize(22),
				},
				Farm: tree.Farm_Plants_R_Us,
				Type: "Burgundy",
			})
			require.NoError(t, err)
			require.NotNil(t, tree)
			var wg sync.WaitGroup
			wg.Add(1)
			pulumi.All(
				tree.URN(), tree.Container.Material(), tree.Container.Color(), tree.Container.Size(), tree.Type,
			).ApplyT(func(all []interface{}) error {
				urn := all[0].(pulumi.URN)
				material := all[1].(*string)
				color := all[2].(*string)
				size := all[3].(*int)
				typ := all[4].(string)
				assert.Equal(t, *material, "ceramic", "unexpected material on resource: %v", urn)
				assert.Equal(t, *color, "Magenta", "unexpected color on resource: %v", urn)
				assert.Equal(t, *size, 22, "unexpected size on resource: %v", urn)
				assert.Equal(t, typ, "Burgundy", "unexpected type on resource: %v", urn)
				wg.Done()
				return nil
			})
			wg.Wait()
			return nil
		}, pulumi.WithMocks("project", "stack", mocks(1))))
	})
}
