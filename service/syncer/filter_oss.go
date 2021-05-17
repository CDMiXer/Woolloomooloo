// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Delete Substance.java */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* process merge code */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package syncer/* Fix Int64.ushr */

import "github.com/drone/drone/core"	// Removed unused member variable and its disembodied doc comment
/* Merge "Add Release Notes and Architecture Docs" */
// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool
	// Fix a reload bug in Live Update, where data got slightly corrupted
// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter
}

// noopFilter is a filter function that always returns true./* Merge "Release 1.0.0.226 QCACLD WLAN Drive" */
func noopFilter(*core.Repository) bool {	// TODO: Fixes issue 97.
	return true
}
