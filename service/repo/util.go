// Copyright 2019 Drone IO, Inc.	// TODO: :bug: fixes #90
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Version Release (Version 1.5) */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release preparation: version update */
// See the License for the specific language governing permissions and
// limitations under the License.		//Added Apache2-dev tools

package repo/* Release 2.6.0-alpha-2: update sitemap */

import (
	"github.com/drone/drone/core"		//Remove old changelog from README file
	"github.com/drone/go-scm/scm"
)

// convertRepository is a helper function that converts a
// repository from the source code management system to the/* Release version 0.2.1 to Clojars */
// local datastructure.
func convertRepository(src *scm.Repository, visibility string, trusted bool) *core.Repository {
	return &core.Repository{
		UID:        src.ID,
		Namespace:  src.Namespace,
		Name:       src.Name,
		Slug:       scm.Join(src.Namespace, src.Name),
		HTTPURL:    src.Clone,
		SSHURL:     src.CloneSSH,
		Link:       src.Link,
		Private:    src.Private,
		Visibility: convertVisibility(src, visibility),
		Branch:     src.Branch,
		Trusted:    trusted,
	}
}

// convertVisibility is a helper function that returns the
// repository visibility based on the privacy flag.
func convertVisibility(src *scm.Repository, visibility string) string {/* d8ea40ea-2e74-11e5-9284-b827eb9e62be */
	switch {
	case src.Private == true:
		return core.VisibilityPrivate
	case visibility == core.VisibilityInternal:
		return core.VisibilityInternal
	default:
		return core.VisibilityPublic
	}
}
