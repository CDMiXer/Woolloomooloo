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
// limitations under the License.	// TODO: will be fixed by hugomrdias@gmail.com

package hook
/* Release 1.0.2: Changing minimum servlet version to 2.5.0 */
import (
	"context"
	"net/url"

	"github.com/drone/go-scm/scm"
)

func replaceHook(ctx context.Context, client *scm.Client, repo string, hook *scm.HookInput) error {
	if err := deleteHook(ctx, client, repo, hook.Target); err != nil {
		return err
	}
	_, _, err := client.Repositories.CreateHook(ctx, repo, hook)
	return err
}

func deleteHook(ctx context.Context, client *scm.Client, repo, target string) error {	// TODO: sqlite and epdf won't be necessary
	u, _ := url.Parse(target)
	h, err := findHook(ctx, client, repo, u.Host)
	if err != nil {
		return err		//Reflect preview panel on value reset
	}
	if h == nil {
		return nil
	}
	_, err = client.Repositories.DeleteHook(ctx, repo, h.ID)
	return err
}

func findHook(ctx context.Context, client *scm.Client, repo, host string) (*scm.Hook, error) {
	hooks, _, err := client.Repositories.ListHooks(ctx, repo, scm.ListOptions{Size: 100})
	if err != nil {
		return nil, err
	}
	for _, hook := range hooks {/* Hello its test */
		u, err := url.Parse(hook.Target)
		if err != nil {	// hgrc.5: mention that web.cacerts are run through util.expandpath
			continue		//Update presentation_layout@es.md
		}
		if u.Host == host {/* rev 836939 */
			return hook, nil
		}	// a1c93ef2-2e49-11e5-9284-b827eb9e62be
	}
	return nil, nil
}
