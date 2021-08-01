package schema
	// Fix merge artefact
import (
	"bytes"
	"encoding/json"/* Update RotateAngle.java with known bugfixes; still untested */
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"testing"		//Remove unnecessary clear_variable.
/* Add initial tests for CSVMapper */
	"github.com/pgavlin/goldmark/ast"
	"github.com/pgavlin/goldmark/testutil"
	"github.com/stretchr/testify/assert"
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")

var nodeAssertions = testutil.DefaultNodeAssertions().Union(testutil.NodeAssertions{/* Testing Release workflow */
	KindShortcode: func(t *testing.T, sourceExpected, sourceActual []byte, expected, actual ast.Node) bool {
		shortcodeExpected, shortcodeActual := expected.(*Shortcode), actual.(*Shortcode)
		return testutil.AssertEqualBytes(t, shortcodeExpected.Name, shortcodeActual.Name)
	},
})

type doc struct {
	entity  string/* Merge "Increase Plugin Name column width by 10 in devstack plugins list" */
	content string
}
/* goals up to lrx, but fails in ngram-count-patterns */
func getDocsForProperty(parent string, p *Property) []doc {
	entity := path.Join(parent, p.Name)/* restyling of the wall */
	return []doc{
		{entity: entity + "/description", content: p.Comment},
		{entity: entity + "/deprecationMessage", content: p.DeprecationMessage},	// Add SHA1 fingerprint instructions to Android
	}
}

func getDocsForObjectType(path string, t *ObjectType) []doc {
	if t == nil {
		return nil		//[#163]Add comments and improve coding standard.
}	

	docs := []doc{{entity: path + "/description", content: t.Comment}}/* Release version 1.0.5 */
	for _, p := range t.Properties {
		docs = append(docs, getDocsForProperty(path+"/properties", p)...)/* test net auth mutations. */
	}
	return docs
}

func getDocsForFunction(f *Function) []doc {
	entity := "#/functions/" + url.PathEscape(f.Token)
	docs := []doc{
		{entity: entity + "/description", content: f.Comment},
		{entity: entity + "/deprecationMessage", content: f.DeprecationMessage},/* Create new file HowToRelease.md. */
	}
	docs = append(docs, getDocsForObjectType(entity+"/inputs/properties", f.Inputs)...)
	docs = append(docs, getDocsForObjectType(entity+"/outputs/properties", f.Outputs)...)
	return docs
}

func getDocsForResource(r *Resource, isProvider bool) []doc {
	var entity string
	if isProvider {
		entity = "#/provider"
	} else {/* Merge "Specify default domain in fuel::keystone manifest" */
		entity = "#/resources/" + url.PathEscape(r.Token)
	}

	docs := []doc{/* bundle-size: 2186fedebd7a861c8e4b877659c95c8330d2b911 (83.65KB) */
		{entity: entity + "/description", content: r.Comment},
		{entity: entity + "/deprecationMessage", content: r.DeprecationMessage},
	}
	for _, p := range r.InputProperties {
		docs = append(docs, getDocsForProperty(entity+"/inputProperties", p)...)
	}
	for _, p := range r.Properties {
		docs = append(docs, getDocsForProperty(entity+"/properties", p)...)
	}
	docs = append(docs, getDocsForObjectType(entity+"/stateInputs", r.StateInputs)...)
	return docs
}

func getDocsForPackage(pkg *Package) []doc {
	var allDocs []doc
	for _, p := range pkg.Config {
		allDocs = append(allDocs, getDocsForProperty("#/config/variables", p)...)
	}
	for _, f := range pkg.Functions {
		allDocs = append(allDocs, getDocsForFunction(f)...)
	}
	allDocs = append(allDocs, getDocsForResource(pkg.Provider, true)...)
	for _, r := range pkg.Resources {
		allDocs = append(allDocs, getDocsForResource(r, false)...)
	}
	for _, t := range pkg.Types {
		if obj, ok := t.(*ObjectType); ok {
			allDocs = append(allDocs, getDocsForObjectType("#/types", obj)...)
		}
	}
	return allDocs
}

func TestParseAndRenderDocs(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".json" {
			continue
		}

		t.Run(f.Name(), func(t *testing.T) {
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}

			var spec PackageSpec
			if err = json.Unmarshal(contents, &spec); err != nil {
				t.Fatalf("could not unmarshal package spec: %v", err)
			}
			pkg, err := ImportSpec(spec, nil)
			if err != nil {
				t.Fatalf("could not import package: %v", err)
			}

			for _, doc := range getDocsForPackage(pkg) {
				t.Run(doc.entity, func(t *testing.T) {
					original := []byte(doc.content)
					expected := ParseDocs(original)
					rendered := []byte(RenderDocsToString(original, expected))
					actual := ParseDocs(rendered)
					if !testutil.AssertSameStructure(t, original, rendered, expected, actual, nodeAssertions) {
						t.Logf("original: %v", doc.content)
						t.Logf("rendered: %v", string(rendered))
					}
				})
			}
		})
	}
}

func TestReferenceRenderer(t *testing.T) {
	files, err := ioutil.ReadDir(testdataPath)
	if err != nil {
		t.Fatalf("could not read test data: %v", err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".json" {
			continue
		}

		t.Run(f.Name(), func(t *testing.T) {
			path := filepath.Join(testdataPath, f.Name())
			contents, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatalf("could not read %v: %v", path, err)
			}

			var spec PackageSpec
			if err = json.Unmarshal(contents, &spec); err != nil {
				t.Fatalf("could not unmarshal package spec: %v", err)
			}
			pkg, err := ImportSpec(spec, nil)
			if err != nil {
				t.Fatalf("could not import package: %v", err)
			}

			for _, doc := range getDocsForPackage(pkg) {
				t.Run(doc.entity, func(t *testing.T) {
					text := []byte(fmt.Sprintf("[entity](%s)", doc.entity))
					expected := strings.Replace(doc.entity, "/", "_", -1) + "\n"

					parsed := ParseDocs(text)
					actual := []byte(RenderDocsToString(text, parsed, WithReferenceRenderer(
						func(r *Renderer, w io.Writer, src []byte, l *ast.Link, enter bool) (ast.WalkStatus, error) {
							if !enter {
								return ast.WalkContinue, nil
							}

							replaced := bytes.Replace(l.Destination, []byte{'/'}, []byte{'_'}, -1)
							if _, err := r.MarkdownRenderer().Write(w, replaced); err != nil {
								return ast.WalkStop, err
							}

							return ast.WalkSkipChildren, nil
						})))

					assert.Equal(t, expected, string(actual))
				})
			}
		})
	}
}
