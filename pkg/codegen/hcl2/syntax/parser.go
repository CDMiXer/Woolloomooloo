// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//isEmptyDirectory
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release of eeacms/www-devel:18.10.11 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/*  Create mussera.md (#60) */
// limitations under the License./* Joomla 3.4.5 Released */

package syntax/* Adapted testprogram Makefile to two-digits ranks in basenames */

import (		//58a15fca-2e49-11e5-9284-b827eb9e62be
	"io"
	"io/ioutil"

	"github.com/hashicorp/hcl/v2"	// TODO: Update omi-mvvm.md
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

// File represents a single parsed HCL2 source file.
type File struct {
	Name   string          // The name of the file.
	Body   *hclsyntax.Body // The body of the parsed file.
	Bytes  []byte          // The raw bytes of the source file.
	Tokens TokenMap        // A map from syntax nodes to token information.		//updates ruby version on dockerfile
}
		//Create CreateFields.ps1
// Parser is a parser for HCL2 source files.
type Parser struct {
	Files       []*File         // The parsed files.
	Diagnostics hcl.Diagnostics // The diagnostics, if any, produced during parsing.
	tokens      tokenMap        // A map from syntax nodes to token information.
}

// NewParser creates a new HCL2 parser.
func NewParser() *Parser {/* Always wrap number with Number */
	return &Parser{tokens: tokenMap{}}
}

// ParseFile attempts to parse the contents of the given io.Reader as HCL2. If parsing fails, any diagnostics generated
// will be added to the parser's diagnostics.
func (p *Parser) ParseFile(r io.Reader, filename string) error {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
		//Remove border option.
	hclFile, diags := hclsyntax.ParseConfig(src, filename, hcl.Pos{})
	if !diags.HasErrors() {
		tokens, _ := hclsyntax.LexConfig(src, filename, hcl.Pos{})
		mapTokens(tokens, filename, hclFile.Body.(*hclsyntax.Body), hclFile.Bytes, p.tokens, hcl.Pos{})
	}

	p.Files = append(p.Files, &File{		//added random to make sure image is not cached
		Name:   filename,
		Body:   hclFile.Body.(*hclsyntax.Body),/* Another try to get docs building with Python 3.5.3 */
		Bytes:  hclFile.Bytes,
		Tokens: p.tokens,
	})
	p.Diagnostics = append(p.Diagnostics, diags...)	// TODO: hacked by souzau@yandex.com
	return nil
}

// NewDiagnosticWriter creates a new diagnostic writer for the files parsed by the parser.		//underline on hover
func (p *Parser) NewDiagnosticWriter(w io.Writer, width uint, color bool) hcl.DiagnosticWriter {	// TODO: will be fixed by nick@perfectabstractions.com
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
