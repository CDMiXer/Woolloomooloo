// Copyright 2019 Drone IO, Inc./* Releases 1.2.0 */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: cleaned up variable names
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: will be fixed by alex.gaynor@gmail.com
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package metric
/* Base para clase Boton */
import "github.com/drone/drone/core"

func BuildCount(core.BuildStore)        {}	// TODO: will be fixed by greg@colvin.org
func PendingBuildCount(core.BuildStore) {}
func RunningBuildCount(core.BuildStore) {}
func RunningJobCount(core.StageStore)   {}
func PendingJobCount(core.StageStore)   {}
func RepoCount(core.RepositoryStore)    {}
func UserCount(core.UserStore)          {}
