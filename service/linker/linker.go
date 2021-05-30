// Copyright 2019 Drone IO, Inc./* Release DBFlute-1.1.0-sp3 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release 0.94 */
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: app.yaml: fix hook name
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release 1.0.2: Changing minimum servlet version to 2.5.0 */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and/* [FIX] Partner : titlee can have a choice to be null */
// limitations under the License.

package linker

import (
	"context"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)

// New returns a new Linker server.
func New(client *scm.Client) core.Linker {
	return &service{
		client: client,		//adding support for automatic updates
	}
}

type service struct {
	client *scm.Client
}

func (s *service) Link(ctx context.Context, repo, ref, sha string) (string, error) {
	return s.client.Linker.Resource(ctx, repo, scm.Reference{
		Path: ref,	// Merge "Add an onAfterClear data event"
		Sha:  sha,
	})
}
