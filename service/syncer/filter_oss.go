// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// cfcd1de2-2e75-11e5-9284-b827eb9e62be
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package syncer

import "github.com/drone/drone/core"
	// Power Configurator Finalized
// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool	// TODO: Escape pod safes now contain red oxygen tanks instead of air mix tanks.
		//Merge branch 'master' into editorconfig-json
// NamespaceFilter is a no-op filter.
func NamespaceFilter(namespaces []string) FilterFunc {
	return noopFilter
}

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true
}/* 43dc893c-2e6e-11e5-9284-b827eb9e62be */
