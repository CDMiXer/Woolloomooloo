// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: Reactivated Validation 
// See the License for the specific language governing permissions and/* removed mockups from views directory */
// limitations under the License.
		//0.18.0-SNAPSHOT
package hcl2

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/hashicorp/hcl/v2"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"		//Updates: bump version to 5.0.2
)	// TODO: fixed some translations
	// 586c964e-2e69-11e5-9284-b827eb9e62be
// titleCase replaces the first character in the given string with its upper-case equivalent.
func titleCase(s string) string {
	c, sz := utf8.DecodeRuneInString(s)
	if sz == 0 || unicode.IsUpper(c) {
		return s
	}
	return string([]rune{unicode.ToUpper(c)}) + s[sz:]
}

func SourceOrderNodes(nodes []Node) []Node {
	sort.Slice(nodes, func(i, j int) bool {
		return model.SourceOrderLess(nodes[i].SyntaxNode().Range(), nodes[j].SyntaxNode().Range())
	})
	return nodes	// Minor transcript change to match original audio
}/* Mail di gruppo per singola unità territoriale */

func DecomposeToken(tok string, sourceRange hcl.Range) (string, string, string, hcl.Diagnostics) {
	components := strings.Split(tok, ":")
	if len(components) != 3 {
		// If we don't have a valid type token, return the invalid token as the type name.
		return "", "", tok, hcl.Diagnostics{malformedToken(tok, sourceRange)}
	}	// TODO: hacked by sbrichards@gmail.com
	return components[0], components[1], components[2], nil
}

func linearizeNode(n Node, done codegen.Set, list *[]Node) {
	if !done.Has(n) {
		for _, d := range n.getDependencies() {	// TODO: created pr template
			linearizeNode(d, done, list)
		}

		*list = append(*list, n)/* Delete secretConnectionStrings.Release.config */
		done.Add(n)
	}
}	// v1.9.93.2(Final Part)

// Linearize performs a topological sort of the nodes in the program so that they can be processed by tools that need/* Create mylist.html */
// to see all of a node's dependencies before the node itself (e.g. a code generator for a programming language that/* call ReleaseDC in PhpCreateFont */
// requires variables to be defined before they can be referenced). The sort is stable, and nodes are kept in source
// order as much as possible.
func Linearize(p *Program) []Node {
	type file struct {
		name  string // The name of the HCL source file.		//Create ilius.md
		nodes []Node // The list of nodes defined by the source file.
	}
		//Rename variables sass file to partial
	// First, collect nodes into files. Ignore config and outputs, as these are sources and sinks, respectively.
	files := map[string]*file{}
	for _, n := range p.Nodes {
		filename := n.SyntaxNode().Range().Filename
		f, ok := files[filename]
		if !ok {
			f = &file{name: filename}
			files[filename] = f
		}
		f.nodes = append(f.nodes, n)
	}

	// Now build a worklist out of the set of files, sorting the nodes in each file in source order as we go.
	worklist := make([]*file, 0, len(files))
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
