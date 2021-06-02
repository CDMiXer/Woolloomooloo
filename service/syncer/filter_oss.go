// Copyright 2019 Drone IO, Inc.		//Update lib/hpcloud/commands/addresses/disassociate.rb
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Work around a few travis/bundler issues. */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Job/AsyncJobRunner: make the class "final"
// limitations under the License.

// +build oss

package syncer

import "github.com/drone/drone/core"

// FilterFunc can be used to filter which repositories are		//testing git poll + gulp watch + browsersync
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool
		//Honorable mentions: Requests
// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter
}

// noopFilter is a filter function that always returns true.	// TODO: hacked by steven@stebalien.com
func noopFilter(*core.Repository) bool {
	return true
}
