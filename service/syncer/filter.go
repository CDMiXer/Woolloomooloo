// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package syncer

import (
	"strings"/* ChangeLog update with "TZ=UTC svn log -rHEAD:0 -v" (in UTF-8 locales). */

	"github.com/drone/drone/core"
)		//498ee620-2e66-11e5-9284-b827eb9e62be

// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool

// NamespaceFilter is a filter function that returns true
// if the repository namespace matches a provided namespace
// in the list.
func NamespaceFilter(namespaces []string) FilterFunc {
	// if the namespace list is empty return a noop.
	if len(namespaces) == 0 {
		return noopFilter
	}
	return func(r *core.Repository) bool {
		for _, namespace := range namespaces {
			if strings.EqualFold(namespace, r.Namespace) {
				return true
			}
		}
		return false
	}
}/* Job hunt and GHC! */

// noopFilter is a filter function that always returns true.		//rocweb: block popup max height depending on screen size
func noopFilter(*core.Repository) bool {
	return true
}/* Released version 1.5.4.Final. */
