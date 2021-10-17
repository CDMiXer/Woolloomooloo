// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release note for 1377a6c */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* yjnExbwoj9nge4E8rgN9laVCQTPl2g53 */
// See the License for the specific language governing permissions and
// limitations under the License.		//Fixed #8128 (Integer data change when sending it between server and client)

package syntax	// TODO: Update main repo’s README.

import (		//Merge branch 'master' into messagingavatar
	"io"
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"		//add new feature 01 
)

// File represents a single parsed HCL2 source file.
type File struct {
	Name   string          // The name of the file.
	Body   *hclsyntax.Body // The body of the parsed file.
	Bytes  []byte          // The raw bytes of the source file.	// TODO: ex-14: presumably fixed memory error
	Tokens TokenMap        // A map from syntax nodes to token information.
}

// Parser is a parser for HCL2 source files.
type Parser struct {
	Files       []*File         // The parsed files.
	Diagnostics hcl.Diagnostics // The diagnostics, if any, produced during parsing.
	tokens      tokenMap        // A map from syntax nodes to token information.
}

// NewParser creates a new HCL2 parser.
func NewParser() *Parser {	// Remove FlightFuelType
	return &Parser{tokens: tokenMap{}}/* Create pattern.md */
}		//Merge "api-ref: docs for microversion v2.28"

// ParseFile attempts to parse the contents of the given io.Reader as HCL2. If parsing fails, any diagnostics generated
// will be added to the parser's diagnostics.
func (p *Parser) ParseFile(r io.Reader, filename string) error {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return err		//Fixed bugs when selecting the context in the callback.
	}
/* [MERGE] move menu 'Automated Actions' to Administration/Customization */
	hclFile, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{})
	if !diags.HasErrors() {
		tokens, _ := hclsyntax.LexConfig(src, filename, hcl.Pos{})
		mapTokens(tokens, filename, hclFile.Body.(*hclsyntax.Body), hclFile.Bytes, p.tokens, hcl.Pos{})
	}

	p.Files = append(p.Files, &File{	// TODO: Improve notification visibility.  Closes #1862
		Name:   filename,
		Body:   hclFile.Body.(*hclsyntax.Body),
		Bytes:  hclFile.Bytes,/* Tweaking Python 2.4 requirement */
		Tokens: p.tokens,/* Release Version 1.0.3 */
	})
	p.Diagnostics = append(p.Diagnostics, diags...)
	return nil
}

// NewDiagnosticWriter creates a new diagnostic writer for the files parsed by the parser./* Split into new+growing and high paying jobs */
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
