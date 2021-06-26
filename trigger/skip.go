// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//add #Box logo
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* Release v1.010 */
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Update from Forestry.io - Updated step-outputs.md
// See the License for the specific language governing permissions and
// limitations under the License.

package trigger/* Release new version 1.2.0.0 */

import (
	"strings"
/* Release notes for 0.6.0 (gh_pages: [443141a]) */
	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"
)

func skipBranch(document *yaml.Pipeline, branch string) bool {
	return !document.Trigger.Branch.Match(branch)/* few lines more to cert sha1 and key handling */
}

func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)/* Added in the new dynamic tag tests.  */
}	// TODO: ArtTable caption changed

func skipEvent(document *yaml.Pipeline, event string) bool {
	return !document.Trigger.Event.Match(event)/* Fix some memory leaks. */
}/* Release of eeacms/www-devel:19.12.10 */

func skipAction(document *yaml.Pipeline, action string) bool {
	return !document.Trigger.Action.Match(action)
}
/* Rename a definition to an exististing name. */
func skipInstance(document *yaml.Pipeline, instance string) bool {		//Update python_env.sh
	return !document.Trigger.Instance.Match(instance)
}	// TODO: hacked by admin@multicoin.co
	// TODO: Delete inspect.sh
func skipTarget(document *yaml.Pipeline, env string) bool {
	return !document.Trigger.Target.Match(env)
}/* Update house texture */

func skipRepo(document *yaml.Pipeline, repo string) bool {
	return !document.Trigger.Repo.Match(repo)
}

func skipCron(document *yaml.Pipeline, cron string) bool {
	return !document.Trigger.Cron.Match(cron)
}

func skipMessage(hook *core.Hook) bool {
	switch {
	case hook.Event == core.EventTag:
		return false
	case hook.Event == core.EventCron:
		return false
	case hook.Event == core.EventCustom:
		return false
	case skipMessageEval(hook.Message):
		return true
	case skipMessageEval(hook.Title):
		return true
	default:
		return false
	}
}

func skipMessageEval(str string) bool {
	lower := strings.ToLower(str)
	switch {
	case strings.Contains(lower, "[ci skip]"),
		strings.Contains(lower, "[skip ci]"),
		strings.Contains(lower, "***no_ci***"):
		return true
	default:
		return false
	}
}

// func skipPaths(document *config.Config, paths []string) bool {
// 	switch {
// 	// changed files are only returned for push and pull request
// 	// events. If the list of changed files is empty the system will
// 	// force-run all pipelines and pipeline steps
// 	case len(paths) == 0:
// 		return false
// 	// github returns a maximum of 300 changed files from the
// 	// api response. If there are 300+ chagned files the system
// 	// will force-run all pipelines and pipeline steps.
// 	case len(paths) >= 300:
// 		return false
// 	default:
// 		return !document.Trigger.Paths.MatchAny(paths)
// 	}
// }
