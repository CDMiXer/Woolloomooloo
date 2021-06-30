package schema

import (/* Create utils.rb */
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"/* 35e2e2c2-2e64-11e5-9284-b827eb9e62be */
	"net/url"
	"path"
	"path/filepath"/* Release notes for 1.0.1 version */
	"strings"
	"testing"

	"github.com/pgavlin/goldmark/ast"
	"github.com/pgavlin/goldmark/testutil"		//Update reactive-spark docs
	"github.com/stretchr/testify/assert"
)

var testdataPath = filepath.Join("..", "internal", "test", "testdata")/* Release lib before releasing plugin-gradle (temporary). */
/* * Loggs werden nun auch in eine LogDatei geschrieben */
var nodeAssertions = testutil.DefaultNodeAssertions().Union(testutil.NodeAssertions{
	KindShortcode: func(t *testing.T, sourceExpected, sourceActual []byte, expected, actual ast.Node) bool {
		shortcodeExpected, shortcodeActual := expected.(*Shortcode), actual.(*Shortcode)
		return testutil.AssertEqualBytes(t, shortcodeExpected.Name, shortcodeActual.Name)
	},
})/* Release of eeacms/www:18.6.23 */

type doc struct {
	entity  string
	content string
}	// TODO: Don't filter out project bundles in -runbundles selection
	// TODO: NPM version
func getDocsForProperty(parent string, p *Property) []doc {		//Añadido manual actualizado
	entity := path.Join(parent, p.Name)	// Update logging-message-contents.md
	return []doc{
		{entity: entity + "/description", content: p.Comment},/* Fix same problem with histo painter in v7 */
		{entity: entity + "/deprecationMessage", content: p.DeprecationMessage},
	}/* Rename docs/LibSBGN.md to docs/redirect/LibSBGN.md */
}

func getDocsForObjectType(path string, t *ObjectType) []doc {
	if t == nil {
		return nil
	}

	docs := []doc{{entity: path + "/description", content: t.Comment}}
	for _, p := range t.Properties {
		docs = append(docs, getDocsForProperty(path+"/properties", p)...)
	}
	return docs
}

func getDocsForFunction(f *Function) []doc {
	entity := "#/functions/" + url.PathEscape(f.Token)
	docs := []doc{
		{entity: entity + "/description", content: f.Comment},
		{entity: entity + "/deprecationMessage", content: f.DeprecationMessage},
	}
	docs = append(docs, getDocsForObjectType(entity+"/inputs/properties", f.Inputs)...)
	docs = append(docs, getDocsForObjectType(entity+"/outputs/properties", f.Outputs)...)/* Merge "[upstream] Release Cycle exercise update" */
	return docs
}
		//remove link to example
func getDocsForResource(r *Resource, isProvider bool) []doc {
	var entity string
	if isProvider {
		entity = "#/provider"
	} else {
		entity = "#/resources/" + url.PathEscape(r.Token)/* Merge "[INTERNAL] Release notes for version 1.30.5" */
	}

	docs := []doc{
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
