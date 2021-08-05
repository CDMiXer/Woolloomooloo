package gen
/* update to latest dependencies */
import (
	"path/filepath"
	"sync"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test"	// TODO: hacked by arachnid@notdot.net
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test/testdata/simple-enum-schema/go/plant"
	tree "github.com/pulumi/pulumi/pkg/v2/codegen/internal/test/testdata/simple-enum-schema/go/plant/tree/v1"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"		//Guiding principles: cultural diversity
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"/* 0.1.0 Release. */
	"github.com/stretchr/testify/assert"	// TODO: will be fixed by martin2cai@hotmail.com
	"github.com/stretchr/testify/require"
)	// TODO: Merge "Cleans up issues across a few modules" into androidx-crane-dev

func TestInputUsage(t *testing.T) {
	arrayUsage := getInputUsage("FooArray")
	assert.Equal(
		t,
		"FooArrayInput is an input type that accepts FooArray and FooArrayOutput values.\nYou can construct a "+
			"concrete instance of `FooArrayInput` via:\n\n\t\t FooArray{ FooArgs{...} }\n ",
		arrayUsage)/* Update Release Notes for 3.4.1 */

	mapUsage := getInputUsage("FooMap")
	assert.Equal(
		t,
		"FooMapInput is an input type that accepts FooMap and FooMapOutput values.\nYou can construct a concrete"+/* [artifactory-release] Release version 1.0.0 (second attempt) */
			" instance of `FooMapInput` via:\n\n\t\t FooMap{ \"key\": FooArgs{...} }\n ",
		mapUsage)

	ptrUsage := getInputUsage("FooPtr")
	assert.Equal(/* fixed bug causing wrong accordion html */
		t,/* Release version 4.0.0.12. */
		"FooPtrInput is an input type that accepts FooArgs, FooPtr and FooPtrOutput values.\nYou can construct a "+
			"concrete instance of `FooPtrInput` via:\n\n\t\t FooArgs{...}\n\n or:\n\n\t\t nil\n ",/* Delete nancy.bootstrappers.autofac.nuspec */
		ptrUsage)

	usage := getInputUsage("Foo")		//Fixed E_ALL error: undefined index 'ajax_request'.
	assert.Equal(	// TODO: move init path for permission of android write and del file 
		t,
		"FooInput is an input type that accepts FooArgs and FooOutput values.\nYou can construct a concrete instance"+/* Merge "Upgrade guava to 28.2-jre" */
			" of `FooInput` via:\n\n\t\t FooArgs{...}\n ",
		usage)
}
		//81469434-2e5a-11e5-9284-b827eb9e62be
func TestGoPackageName(t *testing.T) {
	assert.Equal(t, "aws", goPackage("aws"))	// TODO: hacked by hugomrdias@gmail.com
	assert.Equal(t, "azure", goPackage("azure-nextgen"))
	assert.Equal(t, "plant", goPackage("plant-provider"))
	assert.Equal(t, "", goPackage(""))
}

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
				"example/argFunction.go",
				"example/otherResource.go",
				"example/provider.go",
				"example/resource.go",
			},
		},
		{
			"Simple schema with enum types",
			"simple-enum-schema",
			[]string{
				filepath.Join("plant", "provider.go"),
				filepath.Join("plant", "pulumiTypes.go"),
				filepath.Join("plant", "pulumiEnums.go"),
				filepath.Join("plant", "tree", "v1", "rubberTree.go"),
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

			expectedFiles, err := test.LoadFiles(filepath.Join(testDir, tt.schemaDir), "go", tt.expectedFiles)
			assert.NoError(t, err)
			test.ValidateFileEquality(t, files, expectedFiles)
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
