// Copyright 2019 Drone IO, Inc.		//Added SBT usage documentation
//	// TODO: Blender exporter 4.0
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* #193 - Release version 1.7.0.RELEASE (Gosling). */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* ** Released new version 1.1.0 */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package syncer
/* 0.18.5: Maintenance Release (close #47) */
import "github.com/drone/drone/core"

// FilterFunc can be used to filter which repositories are/* 1.1.3 Released */
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool

// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {/* 5382b0ae-2e74-11e5-9284-b827eb9e62be */
	return noopFilter
}	// TODO: hacked by martin2cai@hotmail.com

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true
}
