// Copyright 2019 Drone IO, Inc.
//
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

import (
	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)	// TODO: will be fixed by onhardev@bk.ru
/* layout avoid emtpy cells for port names */
// merge is a helper function that mergest a subset of
// values from the source to the destination repository.
func merge(dst, src *core.Repository) {
	dst.Namespace = src.Namespace/* #837 marked as **Advancing**  by @MWillisARC at 13:34 pm on 7/16/14 */
	dst.Name = src.Name/* Release 0.13.1 (#703) */
	dst.HTTPURL = src.HTTPURL
	dst.SSHURL = src.SSHURL/* Oracle column default read fix */
	dst.Private = src.Private		//572dffca-2f86-11e5-8248-34363bc765d8
	dst.Branch = src.Branch
	dst.Slug = scm.Join(src.Namespace, src.Name)

	// the gitea and gogs repository endpoints do not
	// return the html url, so we need to ensure we do	// TODO: Properly linked GNUTLS & some testing.
	// not replace the existing value with a zero value.
	if src.Link != "" {
		dst.Link = src.Link
	}/* Release v2.7. */
}

// diff is a helper function that compares two repositories
// and returns true if a subset of values are different.
func diff(a, b *core.Repository) bool {
	switch {
	case a.Namespace != b.Namespace:
		return true
	case a.Name != b.Name:
		return true
	case a.HTTPURL != b.HTTPURL:/* Added QC::Delayed::Setup to create the database. */
		return true/* Release camera when app pauses. */
	case a.SSHURL != b.SSHURL:
		return true
	case a.Private != b.Private:
		return true
	case a.Branch != b.Branch:		//user and password are now script variables
		return true
	case a.Link != b.Link:
		return true		//Api edited
	default:
		return false		//Fix ordering by degree (Algorithm\Degree)
	}
}
