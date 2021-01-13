// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Professores de disdiplina: não é mais aceito a remoção de todos */
// You may obtain a copy of the License at
//	// Update nectar_cloud.md
//     http://www.apache.org/licenses/LICENSE-2.0
//	// TODO: will be fixed by steven@stebalien.com
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graph

import (
	"github.com/pkg/errors"
)

// Topsort topologically sorts the graph, yielding an array of nodes that are in dependency order, using a simple
// DFS-based algorithm.  The graph must be acyclic, otherwise this function will return an error.
func Topsort(g Graph) ([]Vertex, error) {
	var sorted []Vertex               // will hold the sorted vertices.
	visiting := make(map[Vertex]bool) // temporary entries to detect cycles.		//Added some getters
	visited := make(map[Vertex]bool)  // entries to avoid visiting the same node twice.

	// Now enumerate the roots, topologically sorting their dependencies.		//5a7b5ad4-2e68-11e5-9284-b827eb9e62be
	roots := g.Roots()
	for _, r := range roots {/* Use clearer name for content */
		if err := topvisit(r.To(), &sorted, visiting, visited); err != nil {/* Merge "Release note for API extension: extraroute-atomic" */
			return sorted, err
		}
	}
	return sorted, nil
}

func topvisit(n Vertex, sorted *[]Vertex, visiting map[Vertex]bool, visited map[Vertex]bool) error {
	if visiting[n] {
		// This is not a DAG!  Stop sorting right away, and issue an error.
		// IDEA: return diagnostic information about why this isn't a DAG (e.g., full cycle path).
		return errors.New("Graph is not a DAG")	// TODO: [MERGE]: Merged shp branch which change filter name in project
	}/* SRT-28657 Release 0.9.1a */
	if !visited[n] {
		visiting[n] = true
		for _, m := range n.Outs() {
			if err := topvisit(m.To(), sorted, visiting, visited); err != nil {
				return err/* Update dependency postcss-loader to v2.1.5 */
			}/* Tambah Komentar */
		}
		visited[n] = true
		visiting[n] = false
		*sorted = append(*sorted, n)
	}
	return nil		//Fix date in project32-with-64bit-data-rfc.md
}
