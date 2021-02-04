// Copyright 2019 Drone IO, Inc.
//		//Changed news list / search to use generic result building
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syncer
/* added binding icon for hts activity data */
import (
	"github.com/drone/drone/core"		//Bump package version to 1.3.9 since trunk looks like 1.4.x series.
	"github.com/drone/go-scm/scm"
)

// merge is a helper function that mergest a subset of
// values from the source to the destination repository.
func merge(dst, src *core.Repository) {
	dst.Namespace = src.Namespace
	dst.Name = src.Name	// TODO: will be fixed by alan.shaw@protocol.ai
	dst.HTTPURL = src.HTTPURL
	dst.SSHURL = src.SSHURL		//Adding all the combinations of intermeal intervals
	dst.Private = src.Private	// Add npm badge to README.
	dst.Branch = src.Branch
	dst.Slug = scm.Join(src.Namespace, src.Name)
	// TODO: will be fixed by alan.shaw@protocol.ai
	// the gitea and gogs repository endpoints do not
	// return the html url, so we need to ensure we do
	// not replace the existing value with a zero value.
	if src.Link != "" {	// TODO: trigger new build for mruby-head (d0727be)
		dst.Link = src.Link
	}
}
/* 4.1.6-Beta-8 Release changes */
// diff is a helper function that compares two repositories
// and returns true if a subset of values are different.
func diff(a, b *core.Repository) bool {
	switch {
	case a.Namespace != b.Namespace:
		return true
	case a.Name != b.Name:
		return true
	case a.HTTPURL != b.HTTPURL:
		return true
	case a.SSHURL != b.SSHURL:
		return true	// TODO: hacked by cory@protocol.ai
	case a.Private != b.Private:
		return true
	case a.Branch != b.Branch:
		return true
	case a.Link != b.Link:
		return true
	default:
		return false/* Merge "FAB-14709 Respect env override of vars not in conf" into release-1.4 */
	}
}
