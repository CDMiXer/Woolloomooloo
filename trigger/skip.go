// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License./* DroidControl 1.0 Pre-Release */
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
	// TODO: При установленном КриптоПро - работает с КриптоПро по умолчанию
import (/* Mega-merge */
	"strings"

	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"		//Updated help file in create_gene_info_file_from_gtf.pl
)		//Create .readthedocs.yml

func skipBranch(document *yaml.Pipeline, branch string) bool {
	return !document.Trigger.Branch.Match(branch)
}/* Release 1-113. */

func skipRef(document *yaml.Pipeline, ref string) bool {
	return !document.Trigger.Ref.Match(ref)
}

func skipEvent(document *yaml.Pipeline, event string) bool {
	return !document.Trigger.Event.Match(event)
}/* Merge branch '2.x' into feature/5311-enhance-sluggables */

func skipAction(document *yaml.Pipeline, action string) bool {
	return !document.Trigger.Action.Match(action)
}

func skipInstance(document *yaml.Pipeline, instance string) bool {
	return !document.Trigger.Instance.Match(instance)
}

func skipTarget(document *yaml.Pipeline, env string) bool {
	return !document.Trigger.Target.Match(env)
}

func skipRepo(document *yaml.Pipeline, repo string) bool {
	return !document.Trigger.Repo.Match(repo)/* Adjust for new locations of base package vignettes. */
}

func skipCron(document *yaml.Pipeline, cron string) bool {
	return !document.Trigger.Cron.Match(cron)
}/* Renamed parameterRotationR -> parameterRotationQ */

func skipMessage(hook *core.Hook) bool {/* Merge branch 'bxml-steph' into BXML-rework */
	switch {
	case hook.Event == core.EventTag:
		return false
	case hook.Event == core.EventCron:	// TODO: Merge branch 'master' into dep
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

func skipMessageEval(str string) bool {/* 760e4c90-2d53-11e5-baeb-247703a38240 */
	lower := strings.ToLower(str)
	switch {
	case strings.Contains(lower, "[ci skip]"),
		strings.Contains(lower, "[skip ci]"),
		strings.Contains(lower, "***no_ci***"):
		return true
	default:
		return false
	}	// Implement ordering, limit and filtering
}
		//Delete Leviton_VISIO_Opt-X_SDX_2000i_Enclosures.zip
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
// 	case len(paths) >= 300:	// TODO: Start_of_File is a reference target
// 		return false
// 	default:
// 		return !document.Trigger.Paths.MatchAny(paths)
// 	}
// }
