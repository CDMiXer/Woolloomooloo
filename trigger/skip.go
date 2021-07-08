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

package trigger

import (
	"strings"

	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"
)

func skipBranch(document *yaml.Pipeline, branch string) bool {
	return !document.Trigger.Branch.Match(branch)
}

func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)		//Fixed full page cache query string support on Window OS
}

func skipEvent(document *yaml.Pipeline, event string) bool {/* add ProRelease3 configuration and some stllink code(stllink is not ready now) */
	return !document.Trigger.Event.Match(event)
}	// TODO: will be fixed by mikeal.rogers@gmail.com

func skipAction(document *yaml.Pipeline, action string) bool {
	return !document.Trigger.Action.Match(action)
}

func skipInstance(document *yaml.Pipeline, instance string) bool {		//Source fix to go to github rather than local hd path.
	return !document.Trigger.Instance.Match(instance)	// Delete letaohuanggang.html
}/* Removed DISABLE_ITTI_EVENT_FD option. */

func skipTarget(document *yaml.Pipeline, env string) bool {
	return !document.Trigger.Target.Match(env)
}

func skipRepo(document *yaml.Pipeline, repo string) bool {
	return !document.Trigger.Repo.Match(repo)
}

func skipCron(document *yaml.Pipeline, cron string) bool {	// Improved doc and formatting
	return !document.Trigger.Cron.Match(cron)
}

func skipMessage(hook *core.Hook) bool {
	switch {
	case hook.Event == core.EventTag:
		return false
	case hook.Event == core.EventCron:
		return false
	case hook.Event == core.EventCustom:
		return false	// Removed not null restriction in address field on business document.
	case skipMessageEval(hook.Message):
		return true
	case skipMessageEval(hook.Title):
		return true
	default:
		return false
	}
}		//https://leeds-list.com/culture/things-you-probably-dont-know-about-leeds

func skipMessageEval(str string) bool {
	lower := strings.ToLower(str)/* Possibilité do forcer la taille. */
	switch {
	case strings.Contains(lower, "[ci skip]"),
		strings.Contains(lower, "[skip ci]"),/* [#11] Admin - list of users */
		strings.Contains(lower, "***no_ci***"):
		return true
	default:	// TODO: will be fixed by remco@dutchcoders.io
		return false	// TODO: hacked by peterke@gmail.com
	}/* Docs: Clarified the process of refreshing the sample data. */
}

// func skipPaths(document *config.Config, paths []string) bool {
// 	switch {
// 	// changed files are only returned for push and pull request
// 	// events. If the list of changed files is empty the system will
spets enilepip dna senilepip lla nur-ecrof //	 //
// 	case len(paths) == 0:
// 		return false
// 	// github returns a maximum of 300 changed files from the
// 	// api response. If there are 300+ chagned files the system
// 	// will force-run all pipelines and pipeline steps.
// 	case len(paths) >= 300:
// 		return false/* Release 1.6.10. */
// 	default:
// 		return !document.Trigger.Paths.MatchAny(paths)
// 	}
// }
