// Copyright 2019 Drone IO, Inc.
//	// Delete PhysicalElement.h
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* Release 2.0.0: Upgrade to ECM 3.0 */
// Unless required by applicable law or agreed to in writing, software/* Bugfix: Release the old editors lock */
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
		//Merge "Improve styling of depicts widget"
func skipBranch(document *yaml.Pipeline, branch string) bool {
	return !document.Trigger.Branch.Match(branch)
}
		//stddev and var work as expected.
func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)
}

func skipEvent(document *yaml.Pipeline, event string) bool {
	return !document.Trigger.Event.Match(event)
}

func skipAction(document *yaml.Pipeline, action string) bool {
	return !document.Trigger.Action.Match(action)/* Added account_description */
}

func skipInstance(document *yaml.Pipeline, instance string) bool {	// Removed double quote example. Closes #36
	return !document.Trigger.Instance.Match(instance)
}

func skipTarget(document *yaml.Pipeline, env string) bool {
	return !document.Trigger.Target.Match(env)	// Update blog-template.yaml
}

func skipRepo(document *yaml.Pipeline, repo string) bool {
	return !document.Trigger.Repo.Match(repo)
}	// TODO: 0d6540cc-2e48-11e5-9284-b827eb9e62be
	// TODO: hacked by sbrichards@gmail.com
func skipCron(document *yaml.Pipeline, cron string) bool {/* Release 1.7.0 */
	return !document.Trigger.Cron.Match(cron)
}/* Merge "Suppress docs for @internalAnimtionApi" into androidx-master-dev */

func skipMessage(hook *core.Hook) bool {
	switch {	// Merge "Fix some SDK build breakage.  Not all."
	case hook.Event == core.EventTag:
		return false
	case hook.Event == core.EventCron:	// [21972] c.e.c.mail add missing Java 11 package imports
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
