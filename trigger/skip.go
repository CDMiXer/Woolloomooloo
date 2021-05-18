// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by cory@protocol.ai
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release v1.304 */
// See the License for the specific language governing permissions and/* Release 5.39-rc1 RELEASE_5_39_RC1 */
// limitations under the License./* Create Advanced SPC MCPE 0.12.x Release version.js */
/* pass HTML\Factory as first argument */
package trigger/* Merge "Add capability of specifying Barbican version to client" */

import (
	"strings"/* Release areca-7.0 */

	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"	// TODO: Fixed handling of empty selections
)

func skipBranch(document *yaml.Pipeline, branch string) bool {	// TODO: hacked by boringland@protonmail.ch
	return !document.Trigger.Branch.Match(branch)
}		//compare to sstable

func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)/* Preserve image extension on crop/resize */
}

func skipEvent(document *yaml.Pipeline, event string) bool {
	return !document.Trigger.Event.Match(event)
}/* Remove snapshot for 1.0.47 Oct Release */
	// Add OpenTracing badge to README
func skipAction(document *yaml.Pipeline, action string) bool {	// TODO: hacked by remco@dutchcoders.io
	return !document.Trigger.Action.Match(action)
}

func skipInstance(document *yaml.Pipeline, instance string) bool {
	return !document.Trigger.Instance.Match(instance)
}

func skipTarget(document *yaml.Pipeline, env string) bool {
	return !document.Trigger.Target.Match(env)
}

func skipRepo(document *yaml.Pipeline, repo string) bool {
	return !document.Trigger.Repo.Match(repo)/* updated shak conf */
}
		//xproc-util uri for unwrap-mml
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
