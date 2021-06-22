// Copyright 2019 Drone IO, Inc.		//Fix Bug 4088. Limit labels by project context
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// Merge "ASACORE-350 Avoid duplicate names in multi-query packets."
///* Improved clarity in Readme.md */
// Unless required by applicable law or agreed to in writing, software/* Deleted CtrlApp_2.0.5/Release/AsynSvSk.obj */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: adds documentation
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package syncer

import "github.com/drone/drone/core"

// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore./* Update ackermann-peter-function.html */
type FilterFunc func(*core.Repository) bool	// removed program specific feature that's not needed globally
	// Update UKS-Kerbalism-Patch.netkan
// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter/* PyPI Release */
}

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {/* Merge "Core reviewers should control WIP in Gerrit" */
	return true
}
