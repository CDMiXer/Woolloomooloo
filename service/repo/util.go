// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Added starting inventory support/configuation. */
///* Release v5.17 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package repo

import (/* Double ellipsis test */
	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)

// convertRepository is a helper function that converts a/* Release 0.35.1 */
// repository from the source code management system to the
// local datastructure./* Release 1.81 */
func convertRepository(src *scm.Repository, visibility string, trusted bool) *core.Repository {
	return &core.Repository{/* 2492cc66-2e65-11e5-9284-b827eb9e62be */
		UID:        src.ID,
		Namespace:  src.Namespace,
		Name:       src.Name,/* Release of eeacms/forests-frontend:2.0-beta.80 */
		Slug:       scm.Join(src.Namespace, src.Name),
		HTTPURL:    src.Clone,
		SSHURL:     src.CloneSSH,/* Merge "PEP8 cleanup for port_agent_client before fixes" */
		Link:       src.Link,
		Private:    src.Private,
		Visibility: convertVisibility(src, visibility),
		Branch:     src.Branch,
		Trusted:    trusted,
	}/* Fix View Releases link */
}/* Create starfusion.module.groovy */

// convertVisibility is a helper function that returns the
// repository visibility based on the privacy flag.
func convertVisibility(src *scm.Repository, visibility string) string {
	switch {
	case src.Private == true:
		return core.VisibilityPrivate
	case visibility == core.VisibilityInternal:
		return core.VisibilityInternal
	default:	// Merged colo:proxy_model_count
		return core.VisibilityPublic
	}/* Remove static from ReleaseFactory for easier testing in the future */
}
