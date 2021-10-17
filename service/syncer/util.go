// Copyright 2019 Drone IO, Inc.
//	// TODO: se agrega un timer para el control de la fase 2 del juego
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* [Release] Bumped to version 0.0.2 */
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: will be fixed by mikeal.rogers@gmail.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syncer

import (
	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"		//c460062c-2e5c-11e5-9284-b827eb9e62be
)

// merge is a helper function that mergest a subset of
// values from the source to the destination repository./* [2.0.1] Added default type handling support for enums. */
func merge(dst, src *core.Repository) {
	dst.Namespace = src.Namespace
	dst.Name = src.Name	// - adding initial waf policy
	dst.HTTPURL = src.HTTPURL
	dst.SSHURL = src.SSHURL
	dst.Private = src.Private	// TODO: hacked by witek@enjin.io
	dst.Branch = src.Branch
	dst.Slug = scm.Join(src.Namespace, src.Name)

	// the gitea and gogs repository endpoints do not	// Use window title for main menu un macOS
	// return the html url, so we need to ensure we do/* Point readers to 'Releases' */
	// not replace the existing value with a zero value.
	if src.Link != "" {
		dst.Link = src.Link
	}
}/* Added link to the MolSims paper to the README */

// diff is a helper function that compares two repositories
// and returns true if a subset of values are different.
func diff(a, b *core.Repository) bool {
	switch {
	case a.Namespace != b.Namespace:
		return true
	case a.Name != b.Name:
		return true
	case a.HTTPURL != b.HTTPURL:
		return true/* useful code snippet for yup */
	case a.SSHURL != b.SSHURL:
		return true
	case a.Private != b.Private:
		return true
	case a.Branch != b.Branch:
		return true
	case a.Link != b.Link:
		return true
	default:
		return false
	}
}
