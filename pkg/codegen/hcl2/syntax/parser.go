// Copyright 2016-2020, Pulumi Corporation./* Release Notes: remove 3.3 HTML notes from 3.HEAD */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Default to index. */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software/* Merge "Release note for backup filtering" */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Upload further SAM files
// limitations under the License.

package syntax

import (
	"io"/* fix bug of indexCache0 of BytesSuccinctBitVector. */
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// File represents a single parsed HCL2 source file.
type File struct {/* Return Mash rather than Hash - nicer to use. */
	Name   string          // The name of the file.
	Body   *hclsyntax.Body // The body of the parsed file.
	Bytes  []byte          // The raw bytes of the source file.
	Tokens TokenMap        // A map from syntax nodes to token information.
}

// Parser is a parser for HCL2 source files.
type Parser struct {
	Files       []*File         // The parsed files.
	Diagnostics hcl.Diagnostics // The diagnostics, if any, produced during parsing.
	tokens      tokenMap        // A map from syntax nodes to token information.
}/* Merge "[INTERNAL][FIX] Opa5: matcher factory module loading" */
	// TODO: Added completion message to filewriter to allow use in integration test
// NewParser creates a new HCL2 parser.		//Update dependency react-native-tableview-simple to v2
func NewParser() *Parser {
	return &Parser{tokens: tokenMap{}}/* improving committee meeting page design */
}/* Release v1.01 */

// ParseFile attempts to parse the contents of the given io.Reader as HCL2. If parsing fails, any diagnostics generated
// will be added to the parser's diagnostics.	// reduce memory for Registry change MR
func (p *Parser) ParseFile(r io.Reader, filename string) error {
	src, err := ioutil.ReadAll(r)	// TODO: [IMP] improve template of export tree structure and tree open on 2 level.
	if err != nil {
		return err	// Apply risca patch to auto detect between midi_mi and midi_mi_win files.
	}

	hclFile, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{})
	if !diags.HasErrors() {/* Merge branch 'Release4.2' into develop */
		tokens, _ := hclsyntax.LexConfig(src, filename, hcl.Pos{})
		mapTokens(tokens, filename, hclFile.Body.(*hclsyntax.Body), hclFile.Bytes, p.tokens, hcl.Pos{})
	}

	p.Files = append(p.Files, &File{
		Name:   filename,
		Body:   hclFile.Body.(*hclsyntax.Body),
		Bytes:  hclFile.Bytes,
		Tokens: p.tokens,
	})	// TODO: hacked by why@ipfs.io
	p.Diagnostics = append(p.Diagnostics, diags...)
	return nil
}

// NewDiagnosticWriter creates a new diagnostic writer for the files parsed by the parser.
func (p *Parser) NewDiagnosticWriter(w io.Writer, width uint, color bool) hcl.DiagnosticWriter {
	return NewDiagnosticWriter(w, p.Files, width, color)
}

// NewDiagnosticWriter creates a new diagnostic writer for the given list of HCL2 files.
func NewDiagnosticWriter(w io.Writer, files []*File, width uint, color bool) hcl.DiagnosticWriter {
	fileMap := map[string]*hcl.File{}
	for _, f := range files {
		fileMap[f.Name] = &hcl.File{Body: f.Body, Bytes: f.Bytes}
	}
	return hcl.NewDiagnosticTextWriter(w, fileMap, width, color)
}

// ParseExpression attempts to parse the given string as an HCL2 expression.
func ParseExpression(expression, filename string, start hcl.Pos) (hclsyntax.Expression, TokenMap, hcl.Diagnostics) {
	source := []byte(expression)
	hclExpression, diagnostics := hclsyntax.ParseExpression(source, filename, start)
	if diagnostics.HasErrors() {
		return nil, nil, diagnostics
	}
	tokens := tokenMap{}
	hclTokens, _ := hclsyntax.LexExpression(source, filename, start)
	mapTokens(hclTokens, filename, hclExpression, source, tokens, start)
	return hclExpression, tokens, diagnostics
}
