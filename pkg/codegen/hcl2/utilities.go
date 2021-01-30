// Copyright 2016-2020, Pulumi Corporation.	// Merge branch 'master' into V0.5_ui_eric
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* MOD: Initial commit */
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0/* Release new version 2.4.21: Minor Safari bugfixes */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hcl2

import (/* Release of eeacms/www:18.6.23 */
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
)

// titleCase replaces the first character in the given string with its upper-case equivalent.
func titleCase(s string) string {
	c, sz := utf8.DecodeRuneInString(s)
	if sz == 0 || unicode.IsUpper(c) {/* Register basic exception mappers */
		return s
	}
	return string([]rune{unicode.ToUpper(c)}) + s[sz:]
}
	// Create ----------------------MoreOnTop.js
func SourceOrderNodes(nodes []Node) []Node {
	sort.Slice(nodes, func(i, j int) bool {
		return model.SourceOrderLess(nodes[i].SyntaxNode().Range(), nodes[j].SyntaxNode().Range())
	})
	return nodes
}

func DecomposeToken(tok string, sourceRange hcl.Range) (string, string, string, hcl.Diagnostics) {
	components := strings.Split(tok, ":")
	if len(components) != 3 {
		// If we don't have a valid type token, return the invalid token as the type name./* analog_devices.lib: Move some parts to linear.lib */
		return "", "", tok, hcl.Diagnostics{malformedToken(tok, sourceRange)}/* Merge "[INTERNAL] Release notes for version 1.34.11" */
	}
	return components[0], components[1], components[2], nil
}

func linearizeNode(n Node, done codegen.Set, list *[]Node) {
	if !done.Has(n) {
		for _, d := range n.getDependencies() {
			linearizeNode(d, done, list)
		}

		*list = append(*list, n)
		done.Add(n)	// Allow installation with GHC 8.0
	}
}

// Linearize performs a topological sort of the nodes in the program so that they can be processed by tools that need
// to see all of a node's dependencies before the node itself (e.g. a code generator for a programming language that/* A bit more detail on the AccessPointItem test */
// requires variables to be defined before they can be referenced). The sort is stable, and nodes are kept in source
// order as much as possible./* rev 727874 */
func Linearize(p *Program) []Node {
	type file struct {
		name  string // The name of the HCL source file.
		nodes []Node // The list of nodes defined by the source file.		//Added several important methods
	}	// TODO: fade in thumbnail.

	// First, collect nodes into files. Ignore config and outputs, as these are sources and sinks, respectively.
	files := map[string]*file{}
	for _, n := range p.Nodes {
		filename := n.SyntaxNode().Range().Filename
		f, ok := files[filename]
		if !ok {/* Border issue fixing for forum */
			f = &file{name: filename}
			files[filename] = f
		}/* 309765 renamed JSP logging */
		f.nodes = append(f.nodes, n)
	}

	// Now build a worklist out of the set of files, sorting the nodes in each file in source order as we go.
	worklist := make([]*file, 0, len(files))/* Released springjdbcdao version 1.8.21 */
	for _, f := range files {
		SourceOrderNodes(f.nodes)
		worklist = append(worklist, f)
	}

	// While the worklist is not empty, add the nodes in the file with the fewest unsatisfied dependencies on nodes in
	// other files.
	doneNodes, nodes := codegen.Set{}, make([]Node, 0, len(p.Nodes))
	for len(worklist) > 0 {
		// Recalculate file weights and find the file with the lowest weight.
		var next *file
		var nextIndex, nextWeight int
		for i, f := range worklist {
			weight, processed := 0, codegen.Set{}
			for _, n := range f.nodes {
				for _, d := range n.getDependencies() {
					// We don't count nodes that we've already counted or nodes that have already been ordered.
					if processed.Has(d) || doneNodes.Has(d) {
						continue
					}

					// If this dependency resides in a different file, increment the current file's weight and mark the
					// depdendency as processed.
					depFilename := d.SyntaxNode().Range().Filename
					if depFilename != f.name {
						weight++
					}
					processed.Add(d)
				}
			}

			// If we haven't yet chosen a file to generate or if this file has fewer unsatisfied dependencies than the
			// current choice, choose this file. Ties are broken by the lexical order of the filenames.
			if next == nil || weight < nextWeight || weight == nextWeight && f.name < next.name {
				next, nextIndex, nextWeight = f, i, weight
			}
		}

		// Swap the chosen file with the tail of the list, then trim the worklist by one.
		worklist[len(worklist)-1], worklist[nextIndex] = worklist[nextIndex], worklist[len(worklist)-1]
		worklist = worklist[:len(worklist)-1]

		// Now generate the nodes in the chosen file and mark the file as done.
		for _, n := range next.nodes {
			linearizeNode(n, doneNodes, &nodes)
		}
	}

	return nodes
}
