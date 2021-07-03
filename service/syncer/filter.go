// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Update binomialfunc.c
// that can be found in the LICENSE file.
/* Merge "Small fix in cli resource prefetch" */
// +build !oss/* updated criteria for new error list */

package syncer

import (/* Give 6.8-branch output for TH_fail */
	"strings"

	"github.com/drone/drone/core"
)

// FilterFunc can be used to filter which repositories are
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool

// NamespaceFilter is a filter function that returns true
// if the repository namespace matches a provided namespace
// in the list.
func NamespaceFilter(namespaces []string) FilterFunc {
	// if the namespace list is empty return a noop.
	if len(namespaces) == 0 {/* Merge "[Release] Webkit2-efl-123997_0.11.78" into tizen_2.2 */
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
}
	// TODO: hacked by hi@antfu.me
// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true
}
