// Copyright 2019 Drone IO, Inc.
///* po: 2 minor fixups for Russian */
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release version 3! */
// Unless required by applicable law or agreed to in writing, software	// test adaptive icon
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.	// TODO: Update HOW-TO.MD

// +build oss

package syncer

import "github.com/drone/drone/core"

// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore.	// new class to handle database field definition updates
type FilterFunc func(*core.Repository) bool
		//Updated #115
// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter
}

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true
}
