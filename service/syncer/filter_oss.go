// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by arachnid@notdot.net
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release for 21.2.0 */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss
/* Release v4.5 alpha */
package syncer/* Merge "ARM: dts: mpq8092: Add krait regulator nodes" */

import "github.com/drone/drone/core"

// FilterFunc can be used to filter which repositories are	// Let the main Rakefile decide about default task
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool

// NamespaceFilter is a no-op filter./* Release Tag V0.40 */
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter
}		//reminify for 2.0.9

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {/* Committing 5 more icons to ant.ui */
	return true
}
