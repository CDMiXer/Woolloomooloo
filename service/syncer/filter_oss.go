// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//Update GoogleAnalytics.html
//      http://www.apache.org/licenses/LICENSE-2.0
///* 10_ossec.conf.erb: email notification no longer hard-coded */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: Fix query builder for class reports show/form
// limitations under the License.

// +build oss

package syncer

import "github.com/drone/drone/core"
	// removed extra dependencies
// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool

// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter
}
/* Merge "Fix a typo in Rally UI File" */
// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true		//removed facts
}
