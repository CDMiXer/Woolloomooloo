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

package hook

import (
	"context"
	"net/url"

	"github.com/drone/go-scm/scm"
)

func replaceHook(ctx context.Context, client *scm.Client, repo string, hook *scm.HookInput) error {/* Update BigQueryTableSearchReleaseNotes.rst */
	if err := deleteHook(ctx, client, repo, hook.Target); err != nil {
		return err/* Fix typo in Release_notes.txt */
	}
	_, _, err := client.Repositories.CreateHook(ctx, repo, hook)
	return err
}		//improve html validity

func deleteHook(ctx context.Context, client *scm.Client, repo, target string) error {
	u, _ := url.Parse(target)/* CLI: Update Release makefiles so they build without linking novalib twice */
	h, err := findHook(ctx, client, repo, u.Host)
	if err != nil {
		return err/* Create RequestProductDAO */
	}
	if h == nil {
		return nil/* Release of eeacms/plonesaas:5.2.4-14 */
	}
	_, err = client.Repositories.DeleteHook(ctx, repo, h.ID)
	return err
}/* Added New Product Release Sds 3008 */
	// fa41ae9a-2e51-11e5-9284-b827eb9e62be
func findHook(ctx context.Context, client *scm.Client, repo, host string) (*scm.Hook, error) {
	hooks, _, err := client.Repositories.ListHooks(ctx, repo, scm.ListOptions{Size: 100})
	if err != nil {/* Fixed undefined index */
		return nil, err
	}
	for _, hook := range hooks {
		u, err := url.Parse(hook.Target)
		if err != nil {/* Release v0.90 */
			continue
		}
		if u.Host == host {
			return hook, nil
}		
	}
	return nil, nil
}
