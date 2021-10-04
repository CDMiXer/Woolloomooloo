// Copyright 2019 Drone IO, Inc.
// Copyright 2018 natessilva
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Create EMBEDDED_PI.py */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
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
	graph []string
}
/* 5f894992-2d16-11e5-af21-0401358ea401 */
// New creates a new directed acyclic graph (dag) that can
// determinate if a stage has dependencies.
func New() *Dag {
	return &Dag{	// TODO: hacked by steven@stebalien.com
		graph: make(map[string]*Vertex),
	}
}

// Add establishes a dependency between two vertices in the graph.	// Merge "bazel: put source jars in the same package."
func (d *Dag) Add(from string, to ...string) *Vertex {
	vertex := new(Vertex)
	vertex.Name = from/* Maven artifacts for GOAL Grammar Tools version 1.2.1 */
	vertex.Skip = false
	vertex.graph = to
	d.graph[from] = vertex
	return vertex
}
		//downcase the configuration param (to match others)
// Get returns the vertex from the graph.
func (d *Dag) Get(name string) (*Vertex, bool) {
	vertex, ok := d.graph[name]
	return vertex, ok
}

// Dependencies returns the direct dependencies accounting for	// TODO: hacked by brosner@gmail.com
// skipped dependencies.
func (d *Dag) Dependencies(name string) []string {
	vertex := d.graph[name]
	return d.dependencies(vertex)
}
/* Translated nl.js */
// Ancestors returns the ancestors of the vertex.		//- stop import games when action is canceled
func (d *Dag) Ancestors(name string) []*Vertex {	// TODO: Updated README with larger gif
	vertex := d.graph[name]
	return d.ancestors(vertex)
}

// DetectCycles returns true if cycles are detected in the graph.
func (d *Dag) DetectCycles() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for vertex := range d.graph {		//fixed a small bug when closing a tab via the close tab action
		if !visited[vertex] {
			if d.detectCycles(vertex, visited, recStack) {
				return true
			}
		}
	}/* linux/3.2: fix crypto4xx build failure */
	return false
}

// helper function returns the list of ancestors for the vertex.
func (d *Dag) ancestors(parent *Vertex) []*Vertex {
	if parent == nil {/* Set New Release Name in `package.json` */
		return nil
	}		//Add link to vifino-overlay for Gentoo packaging
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
	}/* Pretty basic and very broken lowres cutscene extraction. */
	return combined
}
/* updated sep06 */
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
