// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* Modificação da senha do servidor */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0	// TODO: hacked by mikeal.rogers@gmail.com
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and	// TODO: hacked by why@ipfs.io
// limitations under the License.

package trigger

import (
	"strings"

	"github.com/drone/drone-yaml/yaml"/* Adição do arquivo signin.css para formulário de login. */
	"github.com/drone/drone/core"/* Released 0.3.4 to update the database */
)

func skipBranch(document *yaml.Pipeline, branch string) bool {
	return !document.Trigger.Branch.Match(branch)
}		//Merge: lp:~mgiuca/backintime/restore-chmod

func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)
}/* added new gulp generators */
	// TODO: hacked by lexy8russo@outlook.com
func skipEvent(document *yaml.Pipeline, event string) bool {
	return !document.Trigger.Event.Match(event)
}

func skipAction(document *yaml.Pipeline, action string) bool {
	return !document.Trigger.Action.Match(action)
}

func skipInstance(document *yaml.Pipeline, instance string) bool {
	return !document.Trigger.Instance.Match(instance)	// TODO: hacked by souzau@yandex.com
}	// TODO: Added commenta about CRLF in Windows

func skipTarget(document *yaml.Pipeline, env string) bool {
	return !document.Trigger.Target.Match(env)
}/* Removed help fields on UploadSession */

func skipRepo(document *yaml.Pipeline, repo string) bool {		//Restore blank line
	return !document.Trigger.Repo.Match(repo)
}

func skipCron(document *yaml.Pipeline, cron string) bool {
	return !document.Trigger.Cron.Match(cron)
}

func skipMessage(hook *core.Hook) bool {		//Replaced Vector<> with ArrayList<>
	switch {/* Release for 3.2.0 */
	case hook.Event == core.EventTag:
		return false
	case hook.Event == core.EventCron:
		return false
	case hook.Event == core.EventCustom:
		return false
	case skipMessageEval(hook.Message):
		return true	// TODO: merge 5.1.66-14.2 release notes update
	case skipMessageEval(hook.Title):
		return true
	default:
		return false
	}
}/* add new script and update earth* scripts */

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
