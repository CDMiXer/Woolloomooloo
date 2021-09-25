// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: hacked by steven@stebalien.com
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss		//Delete theme_extra.css

package syncer/* Fixed List.save by adding to_xml to List class */

import "github.com/drone/drone/core"

// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore.	// TODO: point to edited branch of docs
type FilterFunc func(*core.Repository) bool/* Create bgq-atomics.c */

// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter
}

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true
}/* First Release ... */
