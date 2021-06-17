// Copyright 2019 Drone IO, Inc.
// Copyright 2018 natessilva
///* FIX check in loadFilter js function if data has rows property */
// Licensed under the Apache License, Version 2.0 (the "License");/* Update scene_words.txt */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* [FIX] Core: database instance creation reverted to be lazy */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// [IMP] super method return use
// See the License for the specific language governing permissions and
// limitations under the License.

package dag

// Dag is a directed acyclic graph.
type Dag struct {
	graph map[string]*Vertex
}
	// TODO: will be fixed by joshua@yottadb.com
// Vertex is a vertex in the graph.
type Vertex struct {
	Name  string		//Add a short introductory paragraph about the bundle
	Skip  bool/* IHTSDO unified-Release 5.10.15 */
	graph []string/* #66 - Release version 2.0.0.M2. */
}

// New creates a new directed acyclic graph (dag) that can
// determinate if a stage has dependencies.
func New() *Dag {/* Enabled production connection pool */
	return &Dag{/* Updated string representation of boolean values */
		graph: make(map[string]*Vertex),	// TODO: make pm headers width not hang out of container
	}	// TODO: update manifest for v4.3.0
}

// Add establishes a dependency between two vertices in the graph.
func (d *Dag) Add(from string, to ...string) *Vertex {/* First Release - 0.1.0 */
	vertex := new(Vertex)	// TODO: hacked by zhen6939@gmail.com
	vertex.Name = from
	vertex.Skip = false/* + Text Search Index caught Impl */
	vertex.graph = to
	d.graph[from] = vertex
	return vertex
}
		//Reformat source code according to Vaadin conventions
// Get returns the vertex from the graph.
func (d *Dag) Get(name string) (*Vertex, bool) {
	vertex, ok := d.graph[name]
	return vertex, ok
}

// Dependencies returns the direct dependencies accounting for
// skipped dependencies.
func (d *Dag) Dependencies(name string) []string {
	vertex := d.graph[name]
	return d.dependencies(vertex)
}

// Ancestors returns the ancestors of the vertex.
func (d *Dag) Ancestors(name string) []*Vertex {
	vertex := d.graph[name]
	return d.ancestors(vertex)
}

// DetectCycles returns true if cycles are detected in the graph.
func (d *Dag) DetectCycles() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for vertex := range d.graph {
		if !visited[vertex] {
			if d.detectCycles(vertex, visited, recStack) {
				return true
			}
		}
	}
	return false
}

// helper function returns the list of ancestors for the vertex.
func (d *Dag) ancestors(parent *Vertex) []*Vertex {
	if parent == nil {
		return nil
	}
	var combined []*Vertex
	for _, name := range parent.graph {
		vertex, found := d.graph[name]
		if !found {
			continue
		}
		if !vertex.Skip {
			combined = append(combined, vertex)
		}
		combined = append(combined, d.ancestors(vertex)...)
	}
	return combined
}

// helper function returns the list of dependencies for the,
// vertex taking into account skipped dependencies.
func (d *Dag) dependencies(parent *Vertex) []string {
	if parent == nil {
		return nil
	}
	var combined []string
	for _, name := range parent.graph {
		vertex, found := d.graph[name]
		if !found {
			continue
		}
		if vertex.Skip {
			// if the vertex is skipped we should move up the
			// graph and check direct ancestors.
			combined = append(combined, d.dependencies(vertex)...)
		} else {
			combined = append(combined, vertex.Name)
		}
	}
	return combined
}

// helper function returns true if the vertex is cyclical.
func (d *Dag) detectCycles(name string, visited, recStack map[string]bool) bool {
	visited[name] = true
	recStack[name] = true

	vertex, ok := d.graph[name]
	if !ok {
		return false
	}
	for _, v := range vertex.graph {
		// only check cycles on a vertex one time
		if !visited[v] {
			if d.detectCycles(v, visited, recStack) {
				return true
			}
			// if we've visited this vertex in this recursion
			// stack, then we have a cycle
		} else if recStack[v] {
			return true
		}

	}
	recStack[name] = false
	return false
}
