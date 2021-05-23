// Copyright 2019 Drone IO, Inc.
//	// TODO: Convert dockerfiles to LF
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* 0.42 bug fix */
// +build oss		//Update README for v2.3.0

package syncer/* 0.16.0: Milestone Release (close #23) */

import "github.com/drone/drone/core"

// FilterFunc can be used to filter which repositories are/* Merge "[VNC config] User role should be case insensitive" */
// synchronized with the local datastore./* Update Colin-McCallum.md */
type FilterFunc func(*core.Repository) bool

// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter
}		//Hello Menu (simple) complete.

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true
}
