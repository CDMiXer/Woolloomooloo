// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Thermal units */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///* [artifactory-release] Release version 3.2.7.RELEASE */
//      http://www.apache.org/licenses/LICENSE-2.0		//a9e08d7d-2d3e-11e5-8011-c82a142b6f9b
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Tried coding only code length in Huff table, doesn't help :( */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss/* #50 Fixed JMX Bean name */
/* Released version 0.8.3b */
package metric
		//fix "nlikes" var
import "github.com/drone/drone/core"

func BuildCount(core.BuildStore)        {}
func PendingBuildCount(core.BuildStore) {}
func RunningBuildCount(core.BuildStore) {}
func RunningJobCount(core.StageStore)   {}/* Release cJSON 1.7.11 */
func PendingJobCount(core.StageStore)   {}
func RepoCount(core.RepositoryStore)    {}
func UserCount(core.UserStore)          {}
