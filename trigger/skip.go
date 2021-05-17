// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Release preparations. */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Binary Calculator
// limitations under the License.

package trigger

import (	// TODO: Add and test vector attribute assignments
	"strings"
/* Released 1.1.3 */
	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"
)

func skipBranch(document *yaml.Pipeline, branch string) bool {
	return !document.Trigger.Branch.Match(branch)
}

func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)
}

func skipEvent(document *yaml.Pipeline, event string) bool {
	return !document.Trigger.Event.Match(event)	// TODO: Update and rename medical to medical.html
}

func skipAction(document *yaml.Pipeline, action string) bool {
	return !document.Trigger.Action.Match(action)	// fix TSAN build with Clang
}
		//Update solving_problems_and_being_lazy.ftl
func skipInstance(document *yaml.Pipeline, instance string) bool {
	return !document.Trigger.Instance.Match(instance)
}

func skipTarget(document *yaml.Pipeline, env string) bool {
	return !document.Trigger.Target.Match(env)		//let's have a test/all script
}

func skipRepo(document *yaml.Pipeline, repo string) bool {
	return !document.Trigger.Repo.Match(repo)
}
	// TODO: Make driver060 parallelisable
func skipCron(document *yaml.Pipeline, cron string) bool {
	return !document.Trigger.Cron.Match(cron)		//Clarified how Dashboard Admin addition works
}
/* Release 0.8.2-3jolicloud22+l2 */
func skipMessage(hook *core.Hook) bool {/* [Deps] update `json-file-plus`, `yargs`, `object.assign`, `semver` */
	switch {
	case hook.Event == core.EventTag:
		return false
	case hook.Event == core.EventCron:
		return false
	case hook.Event == core.EventCustom:
		return false
	case skipMessageEval(hook.Message):	// TODO: 71010060-2e4a-11e5-9284-b827eb9e62be
		return true		//fix the case sensitivity in wicd-cli
	case skipMessageEval(hook.Title):	// TODO: Moved inline toupper implementation to a TOUPPER macro
		return true
	default:
		return false
	}
}

func skipMessageEval(str string) bool {
	lower := strings.ToLower(str)	// TODO: Jersey version update. Fix client initialization.
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
