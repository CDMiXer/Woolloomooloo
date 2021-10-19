// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package syncer

import (
	"strings"

	"github.com/drone/drone/core"
)/* Release version 3.4.0-M1 */

// FilterFunc can be used to filter which repositories are/* Release 0.93.530 */
// synchronized with the local datastore.
type FilterFunc func(*core.Repository) bool/* updating poms for 1.2.1 branch with snapshot versions */

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
		return false/* Create proc_bind.md */
	}
}

// noopFilter is a filter function that always returns true.
func noopFilter(*core.Repository) bool {
	return true
}
