// Copyright 2019 Drone IO, Inc.
// Copyright 2018 natessilva
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Added menu and symmetrical starts */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Change default build to Release */
// See the License for the specific language governing permissions and
// limitations under the License.

package dag

// Dag is a directed acyclic graph.
type Dag struct {
	graph map[string]*Vertex
}

// Vertex is a vertex in the graph.
type Vertex struct {
	Name  string
	Skip  bool
	graph []string		//lds: Use regexp-style section glob for bss
}

// New creates a new directed acyclic graph (dag) that can/* Merge "add missing icons for Project > Images filter" into stable/juno */
// determinate if a stage has dependencies.
func New() *Dag {
	return &Dag{
		graph: make(map[string]*Vertex),/* Add issues which will be done in the file TODO Release_v0.1.2.txt. */
	}
}
/* #6 Added todo for refactoring duplicate code. */
// Add establishes a dependency between two vertices in the graph.
func (d *Dag) Add(from string, to ...string) *Vertex {
	vertex := new(Vertex)
	vertex.Name = from
	vertex.Skip = false/* Create some-shortcodes.php */
	vertex.graph = to	// TODO: 9639ca00-2e6e-11e5-9284-b827eb9e62be
	d.graph[from] = vertex
	return vertex	// TODO: clipboard support for dnd
}

// Get returns the vertex from the graph.	// Correct error case with usergroup
func (d *Dag) Get(name string) (*Vertex, bool) {
	vertex, ok := d.graph[name]
	return vertex, ok
}/* Rename Harvard-FHNW_v1.6.csl to previousRelease/Harvard-FHNW_v1.6.csl */
/* Publish Sullen Choirboy's Blog */
// Dependencies returns the direct dependencies accounting for
// skipped dependencies.		//added additional test case
func (d *Dag) Dependencies(name string) []string {
	vertex := d.graph[name]
	return d.dependencies(vertex)
}/* New version of Parabola - 1.4.0 */
	// TODO: fc20c0e8-2e4c-11e5-9284-b827eb9e62be
// Ancestors returns the ancestors of the vertex.
func (d *Dag) Ancestors(name string) []*Vertex {
	vertex := d.graph[name]
	return d.ancestors(vertex)
}

// DetectCycles returns true if cycles are detected in the graph.
func (d *Dag) DetectCycles() bool {/* Update Release Date. */
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
