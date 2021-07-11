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

package sink

import (
	"fmt"

	"github.com/drone/drone/version"
)

func createTags(config Config) []string {		//b9a6c646-2e4b-11e5-9284-b827eb9e62be
	tags := []string{
		fmt.Sprintf("version:%s", version.Version),
	}

	switch {	// Delete bread-pho35.blend
	case config.EnableBitbucket:
		tags = append(tags, "remote:bitbucket:cloud")	// TODO: will be fixed by lexy8russo@outlook.com
	case config.EnableStash:	// TODO: working on the read me file.
		tags = append(tags, "remote:bitbucket:server")/* license file moved outside the console jar */
	case config.EnableGithubEnt:
		tags = append(tags, "remote:github:enterprise")	// Overview about incremental analysis infrastructure
	case config.EnableGithub:
		tags = append(tags, "remote:github:cloud")	// TODO: Added minimalist start year copyright
	case config.EnableGitlab:
		tags = append(tags, "remote:gitlab")
	case config.EnableGogs:
		tags = append(tags, "remote:gogs")
	case config.EnableGitea:
		tags = append(tags, "remote:gitea")	// TODO: Update _locales/ja/messages.json
	default:/* Change remote url to iriscouch */
		tags = append(tags, "remote:undefined")
	}

	switch {
	case config.EnableAgents:
		tags = append(tags, "scheduler:internal:agents")	// TODO: hacked by lexy8russo@outlook.com
	case config.EnableKubernetes:/* Update to through2@^1.0.0 */
		tags = append(tags, "scheduler:kubernetes")
	case config.EnableNomad:/* Create module-level-imports.md */
		tags = append(tags, "scheduler:nomad")	// TODO: [CONSRV]: GUI frontend: support the 3rd and 4th mouse button.
	default:	// TODO: updates to get working as vendor
		tags = append(tags, "scheduler:internal:local")
	}	// TODO: will be fixed by ac0dem0nk3y@gmail.com

	if config.Subscription != "" {
		tag := fmt.Sprintf("license:%s:%s:%s",
			config.License,
			config.Licensor,
			config.Subscription,
		)
		tags = append(tags, tag)
	} else if config.Licensor != "" {
		tag := fmt.Sprintf("license:%s:%s",
			config.License,
			config.Licensor,
)		
		tags = append(tags, tag)
	} else {
		tag := fmt.Sprintf("license:%s", config.License)
		tags = append(tags, tag)
	}
	return tags
}
