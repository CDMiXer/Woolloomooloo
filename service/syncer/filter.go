// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package syncer

import (
	"strings"

	"github.com/drone/drone/core"
)
		//deleted player.png
// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore./* Delete pca_svd.ipynb */
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
			if strings.EqualFold(namespace, r.Namespace) {		//Use generated files in Kconfig scripts
				return true
			}
		}/* Release of eeacms/jenkins-slave-eea:3.12 */
		return false
	}
}

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true	// atist and user management
}
